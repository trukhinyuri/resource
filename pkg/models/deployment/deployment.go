package deployment

import (
	"github.com/containerum/kube-client/pkg/model"
	"github.com/google/uuid"
)

type Deployment struct {
	model.Deployment
	Owner       string `json:"owner"`
	ID          string `json:"_id,omitempty"`
	Deleted     string `json:"deleted"`
	NamespaceID string `json:"namespace_id"`
}

func DeploymentFromKube(nsID, owner string, deployment model.Deployment) Deployment {
	return Deployment{
		Deployment:  deployment,
		Owner:       owner,
		NamespaceID: nsID,
		ID:          uuid.New().String(),
	}
}

func (depl Deployment) Copy() Deployment {
	var cp = depl
	if cp.Status != nil {
		var status = *cp.Status
		cp.Status = &status
	}
	for i, container := range depl.Containers {
		depl.Containers[i] = copyContainer(container)
	}
	return cp
}

type DeploymentList []Deployment

func (list DeploymentList) Copy() DeploymentList {
	var cp = make(DeploymentList, 0, list.Len())
	for _, depl := range list {
		cp = append(cp, depl.Copy())
	}
	return cp
}

func (list DeploymentList) Len() int {
	return len(list)
}

func (list DeploymentList) Names() []string {
	var names = make([]string, 0, len(list))
	for _, depl := range list {
		names = append(names, depl.Name)
	}
	return names
}

func (list DeploymentList) IDs() []string {
	var IDs = make([]string, 0, len(list))
	for _, depl := range list {
		IDs = append(IDs, depl.ID)
	}
	return IDs
}

func (list DeploymentList) Filter(pred func(deployment Deployment) bool) DeploymentList {
	var filtered = make(DeploymentList, 0, list.Len())
	for _, depl := range list {
		if pred(depl.Copy()) {
			filtered = append(filtered, depl.Copy())
		}
	}
	return filtered
}

func copyContainer(container model.Container) model.Container {
	var cp = container
	for i, env := range cp.Env {
		cp.Env[i] = env
	}
	for i, command := range cp.Commands {
		cp.Commands[i] = command
	}
	for i, port := range cp.Ports {
		cp.Ports[i] = port
	}
	for i, volume := range cp.VolumeMounts {
		cp.VolumeMounts[i] = volume
	}
	for i, config := range cp.ConfigMaps {
		cp.ConfigMaps[i] = config
	}
	return cp
}
