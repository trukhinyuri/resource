package server

import (
	"errors"
	"io"
	"reflect"
	"sync"

	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"context"

	"git.containerum.net/ch/json-types/billing"
	rstypes "git.containerum.net/ch/json-types/resource-service"
	"git.containerum.net/ch/kube-client/pkg/cherry/resource-service"
	kubtypes "git.containerum.net/ch/kube-client/pkg/model"
	"git.containerum.net/ch/resource-service/pkg/models"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Parallel runs functions in dedicated goroutines and waits for ending
func Parallel(funcs ...func() error) (ret []error) {
	wg := &sync.WaitGroup{}
	wg.Add(len(funcs))
	retmu := &sync.Mutex{}
	for _, f := range funcs {
		go func(inf func() error) {
			if err := inf(); err != nil {
				retmu.Lock()
				defer retmu.Unlock()
				ret = append(ret, err)
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	return
}

// CheckTariff checks if user has permissions to use tariff
func CheckTariff(tariff billing.Tariff, isAdmin bool) error {
	if !tariff.Active {
		return rserrors.ErrTariffNotFound()
	}
	if !isAdmin && !tariff.Public {
		return rserrors.ErrTariffNotFound()
	}

	return nil
}

// VolumeLabel generates label for non-persistent volume
func VolumeLabel(nsLabel string) string {
	return nsLabel + "-volume"
}

// DetermineServiceType deduces service type from service ports. If we have one or more "Port" set it is internal.
func DetermineServiceType(service kubtypes.Service) rstypes.ServiceType {
	serviceType := rstypes.ServiceExternal
	for _, port := range service.Ports {
		if port.Port != nil {
			serviceType = rstypes.ServiceInternal
			break
		}
	}
	return serviceType
}

// IngressPaths generates ingress paths by service ports
func IngressPaths(service kubtypes.Service, path string, servicePort int) ([]kubtypes.Path, error) {
	var portExist bool
	for _, port := range service.Ports {
		if port.Port != nil && *port.Port == servicePort && port.Protocol == kubtypes.TCP {
			portExist = true
			break
		}
	}
	if !portExist {
		return nil, rserrors.ErrTCPPortNotFound().AddDetailF("TCP port %d not exists in service %s", servicePort, service.Name)
	}

	ret := []kubtypes.Path{
		{Path: path, ServiceName: service.Name, ServicePort: servicePort},
	}

	return ret, nil
}

// VolumeGlusterName generates volume name for glusterfs (non-persistent volumes)
func VolumeGlusterName(nsLabel, userID string) string {
	glusterName := sha256.Sum256([]byte(fmt.Sprintf("%s-volume%s", nsLabel, userID)))
	return hex.EncodeToString(glusterName[:])
}

func GetAndCheckPermission(ctx context.Context, db models.AccessDB, userID string, resourceKind rstypes.Kind, resourceName string, needed rstypes.PermissionStatus) error {
	if IsAdminRole(ctx) {
		return nil
	}

	current, err := db.GetUserResourceAccess(ctx, userID, resourceKind, resourceName)
	if err != nil {
		return err
	}

	if !models.PermCheck(current, needed) {
		return rserrors.ErrPermissionDenied().AddDetailF("permission '%s' required for operation, you have '%s'", needed, current)
	}

	return nil
}

func CheckNamespaceResize(ns models.NamespaceUsage, newTariff billing.NamespaceTariff) error {
	if newTariff.CPULimit < ns.CPU ||
		newTariff.MemoryLimit < ns.RAM ||
		newTariff.ExternalServices < ns.ExtServices ||
		newTariff.InternalServices < ns.IntServices {
		return rserrors.ErrDownResizeNotAllowed()
	}
	return nil
}

func CheckDeploymentCreateQuotas(ns rstypes.Namespace, nsUsage models.NamespaceUsage, deploy kubtypes.Deployment) error {
	if err := CalculateDeployResources(&deploy); err != nil {
		return err
	}

	var deployCPU, deployRAM int
	cpuQty, err := resource.ParseQuantity(deploy.TotalCPU)
	if err != nil {
		return rserrors.ErrInternal().AddDetailF("unable to parse CPU quota").AddDetailsErr(err)
	}
	ramQty, err := resource.ParseQuantity(deploy.TotalMemory)
	if err != nil {
		return rserrors.ErrInternal().AddDetailF("unable to parse Memory quota").AddDetailsErr(err)
	}
	deployCPU = int(cpuQty.ScaledValue(resource.Milli))
	deployRAM = int(ramQty.ScaledValue(resource.Mega))

	if exceededCPU := ns.CPU - deployCPU - nsUsage.CPU; exceededCPU < 0 {
		return rserrors.ErrQuotaExceeded().AddDetailF("Exceeded %d CPU", -exceededCPU)
	}

	if exceededRAM := ns.RAM - deployRAM - nsUsage.RAM; exceededRAM < 0 {
		return rserrors.ErrQuotaExceeded().AddDetailF("Exceeded %d memory", -exceededRAM)
	}

	return nil
}

func CheckDeploymentReplaceQuotas(ns rstypes.Namespace, nsUsage models.NamespaceUsage, oldDeploy, newDeploy kubtypes.Deployment) error {
	if err := CalculateDeployResources(&oldDeploy); err != nil {
		return err
	}

	var oldDeployCPU, oldDeployRAM int
	cpuQty, _ := resource.ParseQuantity(oldDeploy.TotalCPU)
	ramQty, _ := resource.ParseQuantity(oldDeploy.TotalMemory)
	oldDeployCPU = int(cpuQty.ScaledValue(resource.Milli))
	oldDeployRAM = int(ramQty.ScaledValue(resource.Mega))

	if err := CalculateDeployResources(&newDeploy); err != nil {
		return err
	}

	var newDeployCPU, newDeployRAM int
	cpuQty, _ = resource.ParseQuantity(newDeploy.TotalCPU)
	ramQty, _ = resource.ParseQuantity(newDeploy.TotalMemory)
	newDeployCPU = int(cpuQty.ScaledValue(resource.Milli))
	newDeployRAM = int(ramQty.ScaledValue(resource.Mega))

	if exceededCPU := ns.CPU - nsUsage.CPU - newDeployCPU + oldDeployCPU; exceededCPU < 0 {
		return rserrors.ErrQuotaExceeded().AddDetailF("Exceeded %d CPU", -exceededCPU)
	}

	if exceededRAM := ns.CPU - nsUsage.CPU - newDeployRAM + oldDeployRAM; exceededRAM < 0 {
		return rserrors.ErrQuotaExceeded().AddDetailF("Exceeded %d memory", -exceededRAM)
	}

	return nil
}

func CheckDeploymentReplicasChangeQuotas(ns rstypes.Namespace, nsUsage models.NamespaceUsage, deploy kubtypes.Deployment, newReplicas int) error {
	if err := CalculateDeployResources(&deploy); err != nil {
		return err
	}

	var deployCPU, deployRAM int
	cpuQty, _ := resource.ParseQuantity(deploy.TotalCPU)
	ramQty, _ := resource.ParseQuantity(deploy.TotalMemory)
	deployCPU = int(cpuQty.ScaledValue(resource.Milli))
	deployRAM = int(ramQty.ScaledValue(resource.Mega))

	if exceededCPU := ns.CPU - nsUsage.CPU - deployCPU*newReplicas + deployCPU*deploy.Replicas; exceededCPU < 0 {
		return rserrors.ErrQuotaExceeded().AddDetailF("Exceeded %d CPU", -exceededCPU)
	}

	if exceededRAM := ns.CPU - nsUsage.CPU - deployRAM*newReplicas + deployRAM*deploy.Replicas; exceededRAM < 0 {
		return rserrors.ErrQuotaExceeded().AddDetailF("Exceeded %d memory", -exceededRAM)
	}

	return nil
}

func CheckServiceCreateQuotas(ns rstypes.Namespace, nsUsage models.NamespaceUsage, serviceType rstypes.ServiceType) error {
	switch serviceType {
	case rstypes.ServiceExternal:
		if ns.MaxExternalServices >= nsUsage.ExtServices {
			return rserrors.ErrQuotaExceeded().AddDetailF("Maximum of external services reached")
		}
	case rstypes.ServiceInternal:
		if ns.MaxIntServices >= nsUsage.IntServices {
			return rserrors.ErrQuotaExceeded().AddDetailF("Maximum of internal services reached")
		}
	default:
		return rserrors.ErrValidation().AddDetailF("Invalid service type %s", serviceType)
	}
	return nil
}

func CalculateDeployResources(deploy *kubtypes.Deployment) error {
	var mCPU, mbRAM int64
	for i, container := range deploy.Containers {
		containerCPU, err := resource.ParseQuantity(container.Limits.CPU)
		if err != nil {
			return rserrors.ErrValidation().AddDetailF("Container %d CPU limit parse failed", i)
		}
		containerRAM, err := resource.ParseQuantity(container.Limits.Memory)
		if err != nil {
			return rserrors.ErrValidation().AddDetailF("Container %d memory limit parse failed", 1)
		}
		mCPU += containerCPU.ScaledValue(resource.Milli)
		mbRAM += containerRAM.ScaledValue(resource.Mega)
	}
	mCPU *= int64(deploy.Replicas)
	mbRAM *= int64(deploy.Replicas)
	deploy.TotalCPU = resource.NewScaledQuantity(mCPU, resource.Milli).String()
	deploy.TotalMemory = resource.NewScaledQuantity(mbRAM, resource.Mega).String()
	return nil
}

func (rs *ResourceServiceClients) UpdateAccess(ctx context.Context, db models.AccessDB, userID string) error {
	accesses, err := db.GetUserResourceAccesses(ctx, userID)
	if err != nil {
		return err
	}
	return rs.Auth.UpdateUserAccess(ctx, userID, accesses)
}

func (rs *ResourceServiceClients) Close() error {
	var errs []string
	v := reflect.ValueOf(rs)
	for i := 0; i < v.NumField(); i++ {
		if closer, ok := v.Field(i).Interface().(io.Closer); ok {
			if err := closer.Close(); err != nil {
				errs = append(errs, closer.Close().Error())
			}
		}
	}
	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("%#v", errs))
	}
	return nil
}
