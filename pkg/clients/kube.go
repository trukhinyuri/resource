package clients

import (
	"context"
	"fmt"
	"net/url"

	"git.containerum.net/ch/resource-service/pkg/rserrors"
	"git.containerum.net/ch/resource-service/pkg/util/coblog"
	"github.com/containerum/cherry"
	"github.com/containerum/cherry/adaptors/cherrylog"
	kubtypes "github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/utils/httputil"
	"github.com/go-resty/resty"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

// Kube is an interface to kube-api service
type Kube interface {
	GetDeployment(ctx context.Context, nsID, deployName string) (*kubtypes.Deployment, error)
	CreateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) error
	UpdateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) error
	SetDeploymentReplicas(ctx context.Context, nsID, deplName string, replicas int) error
	SetContainerImage(ctx context.Context, nsID, deplName string, container kubtypes.UpdateImage) error
	DeleteSolutionDeployments(ctx context.Context, nsID, solutionName string) error
	DeleteDeployment(ctx context.Context, nsID, deplName string) error

	CreateIngress(ctx context.Context, nsID string, ingress kubtypes.Ingress) error
	UpdateIngress(ctx context.Context, nsID string, ingress kubtypes.Ingress) error
	DeleteIngress(ctx context.Context, nsID, ingressName string) error

	CreateSecret(ctx context.Context, nsID string, secret kubtypes.Secret) error
	DeleteSecret(ctx context.Context, nsID, secretName string) error

	GetService(ctx context.Context, nsID, svcName string) (*kubtypes.Service, error)
	CreateService(ctx context.Context, nsID string, service kubtypes.Service) error
	UpdateService(ctx context.Context, nsID string, service kubtypes.Service) error
	DeleteService(ctx context.Context, nsID, serviceName string) error
	DeleteSolutionServices(ctx context.Context, nsID, solutionName string) error

	CreateConfigMap(ctx context.Context, nsID string, cm kubtypes.ConfigMap) error
	DeleteConfigMap(ctx context.Context, nsID, cmName string) error
}

type kube struct {
	client *resty.Client
	log    *cherrylog.LogrusAdapter
}

// NewKubeHTTP creates http client to kube-api service.
func NewKubeHTTP(u *url.URL) Kube {
	log := logrus.WithField("component", "kube_client")
	client := resty.New().
		SetHostURL(u.String()).
		SetLogger(log.WriterLevel(logrus.DebugLevel)).
		SetDebug(true).
		SetError(cherry.Err{}).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	return kube{
		client: client,
		log:    cherrylog.NewLogrusAdapter(log),
	}
}

func (kub kube) GetDeployment(ctx context.Context, nsID, deployName string) (*kubtypes.Deployment, error) {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("get deployment %v", deployName)

	var ret kubtypes.Deployment
	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetResult(&ret).
		SetPathParams(map[string]string{
			"namespace":  nsID,
			"deployment": deployName,
		}).
		Get("/namespaces/{namespace}/deployments/{deployment}")
	if err != nil {
		return nil, rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return nil, resp.Error().(*cherry.Err)
	}
	return &ret, nil
}

func (kub kube) CreateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) error {
	kub.log.WithField("ns_id", nsID).Debugf("create deployment %v", deploy.Name)
	coblog.Std.Struct(deploy)

	resp, err := kub.client.R().
		SetBody(deploy).
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetPathParams(map[string]string{
			"namespace": nsID,
		}).
		Post("/namespaces/{namespace}/deployments")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) DeleteDeployment(ctx context.Context, nsID, deployName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"deploy_name": deployName,
	}).Debug("delete deployment")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetPathParams(map[string]string{
			"namespace":  nsID,
			"deployment": deployName,
		}).
		Delete("/namespaces/{namespace}/deployments/{deployment}")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) DeleteSolutionDeployments(ctx context.Context, nsID, solutionName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":    nsID,
		"solution": solutionName,
	}).Debug("delete solution deployments")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"solution":  solutionName,
		}).
		Delete("/namespaces/{namespace}/solutions/{solution}/deployments")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) UpdateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("update deployment %v", deploy.Name)
	coblog.Std.Struct(deploy)

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(deploy).
		SetPathParams(map[string]string{
			"namespace":  nsID,
			"deployment": deploy.Name,
		}).
		Put("/namespaces/{namespace}/deployments/{deployment}")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) SetDeploymentReplicas(ctx context.Context, nsID, deplName string, replicas int) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"deploy_name": deplName,
		"replicas":    replicas,
	}).Debug("change replicas")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(kubtypes.UpdateReplicas{Replicas: replicas}).
		SetPathParams(map[string]string{
			"namespace":  nsID,
			"deployment": deplName,
		}).
		Put("/namespaces/{namespace}/deployments/{deployment}/replicas")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) SetContainerImage(ctx context.Context, nsID, deplName string, container kubtypes.UpdateImage) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"deploy_name": deplName,
		"container":   container.Container,
		"image":       container.Image,
	}).Debug("set container image")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(container).
		SetPathParams(map[string]string{
			"namespace":  nsID,
			"deployment": deplName,
		}).
		Put("/namespaces/{namespace}/deployments/{deployment}/image")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) CreateIngress(ctx context.Context, nsID string, ingress kubtypes.Ingress) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("create ingress %v", ingress.Name)
	coblog.Std.Struct(ingress)

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(ingress).
		SetPathParams(map[string]string{
			"namespace": nsID,
		}).
		Post("/namespaces/{namespace}/ingresses")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) UpdateIngress(ctx context.Context, nsID string, ingress kubtypes.Ingress) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":        nsID,
		"ingress_name": ingress.Name,
	}).Debugf("update ingress to %v", ingress.Name)
	coblog.Std.Struct(ingress)

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(ingress).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"ingress":   ingress.Name,
		}).
		Put("/namespaces/{namespace}/ingresses/{ingress}")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) DeleteIngress(ctx context.Context, nsID, ingressName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":        nsID,
		"ingress_name": ingressName,
	}).Debug("delete ingress")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"ingress":   ingressName,
		}).
		Delete("/namespaces/{namespace}/ingresses/{ingress}")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) CreateSecret(ctx context.Context, nsID string, secret kubtypes.Secret) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("create secret %v", secret.Name)
	coblog.Std.Struct(secret)

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(secret).
		SetPathParams(map[string]string{
			"namespace": nsID,
		}).
		Post("/namespaces/{namespace}/secrets")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) DeleteSecret(ctx context.Context, nsID, secretName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"secret_name": secretName,
	}).Debug("delete secret")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"secret":    secretName,
		}).
		Delete("/namespaces/{namespace}/secrets/{secret}")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}

	return nil
}

func (kub kube) GetService(ctx context.Context, nsID, svcName string) (*kubtypes.Service, error) {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("get service %v", svcName)

	var ret kubtypes.Service
	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetResult(&ret).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"service":   svcName,
		}).
		Get("/namespaces/{namespace}/services/{service}")
	if err != nil {
		return nil, rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return nil, resp.Error().(*cherry.Err)
	}
	return &ret, nil
}

func (kub kube) CreateService(ctx context.Context, nsID string, service kubtypes.Service) error {
	kub.log.WithField("ns_id", nsID).Debugf("create service %v", service)
	coblog.Std.Struct(service)

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(service).
		SetPathParams(map[string]string{
			"namespace": nsID,
		}).
		Post("/namespaces/{namespace}/services")

	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}

	return nil
}

func (kub kube) UpdateService(ctx context.Context, nsID string, service kubtypes.Service) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":        nsID,
		"service_name": service.Name,
	}).Debugf("update service to %v", service.Name)
	coblog.Std.Struct(service)

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(service).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"service":   service.Name,
		}).
		Put("/namespaces/{namespace}/services/{service}")

	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}

	return nil
}

func (kub kube) DeleteService(ctx context.Context, nsID, serviceName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":        nsID,
		"service_name": serviceName,
	}).Debug("delete service")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"service":   serviceName,
		}).
		Delete("/namespaces/{namespace}/services/{service}")

	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}

	return nil
}

func (kub kube) DeleteSolutionServices(ctx context.Context, nsID, solutionName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":    nsID,
		"solution": solutionName,
	}).Debug("delete solution services")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"solution":  solutionName,
		}).
		Delete("/namespaces/{namespace}/solutions/{solution}/services")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) CreateConfigMap(ctx context.Context, nsID string, cm kubtypes.ConfigMap) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("create configmap %v", cm.Name)
	coblog.Std.Struct(cm)

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetBody(cm).
		SetPathParams(map[string]string{
			"namespace": nsID,
		}).
		Post("/namespaces/{namespace}/configmaps")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (kub kube) DeleteConfigMap(ctx context.Context, nsID, cmName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":   nsID,
		"cm_name": cmName,
	}).Debug("delete configmap")

	resp, err := kub.client.R().
		SetContext(ctx).
		SetHeaders(httputil.RequestXHeadersMap(ctx)).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"configmap": cmName,
		}).
		SetPathParams(map[string]string{
			"namespace": nsID,
			"configmap": cmName,
		}).
		Delete("/namespaces/{namespace}/configmaps/{configmap}")
	if err != nil {
		return rserrors.ErrInternal().Log(err, kub.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}

	return nil
}

func (kub kube) String() string {
	return fmt.Sprintf("kube api http client: url=%v", kub.client.HostURL)
}

// Dummy implementation

type kubeDummy struct {
	log *logrus.Entry
}

// NewDummyKube creates a dummy client to kube-api service. It does nothing but logs actions.
func NewDummyKube() Kube {
	return kubeDummy{log: logrus.WithField("component", "kube_stub")}
}

func (kub kubeDummy) GetDeployment(ctx context.Context, nsID, deployName string) (*kubtypes.Deployment, error) {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("get deployment %v", deployName)

	return nil, nil
}

func (kub kubeDummy) CreateDeployment(_ context.Context, nsID string, deploy kubtypes.Deployment) error {
	kub.log.WithField("ns_id", nsID).Debug("create deployment %+v", deploy)

	return nil
}

func (kub kubeDummy) UpdateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("update deployment %+v", deploy)

	return nil
}

func (kub kubeDummy) DeleteDeployment(ctx context.Context, nsID, deplName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"deploy_name": deplName,
	}).Debug("delete deployment")

	return nil
}

func (kub kubeDummy) DeleteSolutionDeployments(ctx context.Context, nsID, solutionName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":    nsID,
		"solution": solutionName,
	}).Debug("delete solution deployments")

	return nil
}

func (kub kubeDummy) ReplaceDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"deploy_name": deploy.Name,
	}).Debugf("replace deployment %+v", deploy)

	return nil
}

func (kub kubeDummy) SetDeploymentReplicas(ctx context.Context, nsID, deplName string, replicas int) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"deploy_name": deplName,
		"replicas":    replicas,
	}).Debug("change replicas")

	return nil
}

func (kub kubeDummy) SetContainerImage(ctx context.Context, nsID, deplName string, container kubtypes.UpdateImage) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"deploy_name": deplName,
		"container":   container.Container,
		"image":       container.Image,
	}).Debug("set container image")

	return nil
}

func (kub kubeDummy) CreateIngress(ctx context.Context, nsID string, ingress kubtypes.Ingress) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("create ingress %+v", ingress)

	return nil
}

func (kub kubeDummy) UpdateIngress(ctx context.Context, nsID string, ingress kubtypes.Ingress) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":        nsID,
		"ingress_name": ingress.Name,
	}).Debugf("update ingress to %+v", ingress)

	return nil
}

func (kub kubeDummy) DeleteIngress(ctx context.Context, nsID, ingressName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":        nsID,
		"ingress_name": ingressName,
	}).Debug("delete ingress")

	return nil
}

func (kub kubeDummy) CreateSecret(ctx context.Context, nsID string, secret kubtypes.Secret) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("create secret %+v", secret)

	return nil
}

func (kub kubeDummy) DeleteSecret(ctx context.Context, nsID, secretName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":       nsID,
		"secret_name": secretName,
	}).Debug("delete secret")

	return nil
}

func (kub kubeDummy) GetService(ctx context.Context, nsID, svcName string) (*kubtypes.Service, error) {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("get service %v", svcName)

	return nil, nil
}

func (kub kubeDummy) CreateService(ctx context.Context, nsID string, service kubtypes.Service) error {
	kub.log.WithField("ns_id", nsID).Debugf("create service %+v", service)

	return nil
}

func (kub kubeDummy) UpdateService(ctx context.Context, nsID string, service kubtypes.Service) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":        nsID,
		"service_name": service.Name,
	}).Debugf("update service to %+v", service)

	return nil
}

func (kub kubeDummy) DeleteService(ctx context.Context, nsID, serviceName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":        nsID,
		"service_name": serviceName,
	}).Debug("delete service")

	return nil
}

func (kub kubeDummy) DeleteSolutionServices(ctx context.Context, nsID, solutionName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":    nsID,
		"solution": solutionName,
	}).Debug("delete solution services")

	return nil
}

func (kub kubeDummy) CreateConfigMap(ctx context.Context, nsID string, cm kubtypes.ConfigMap) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id": nsID,
	}).Debugf("create configmap %v", cm.Name)

	return nil
}

func (kub kubeDummy) DeleteConfigMap(ctx context.Context, nsID, cmName string) error {
	kub.log.WithFields(logrus.Fields{
		"ns_id":   nsID,
		"cm_name": cmName,
	}).Debug("delete configmap")

	return nil
}

func (kubeDummy) String() string {
	return "kube api dummy"
}
