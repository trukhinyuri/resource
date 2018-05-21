package server

import (
	"context"

	"git.containerum.net/ch/resource-service/pkg/models/deployment"
	"git.containerum.net/ch/resource-service/pkg/models/domain"
	"git.containerum.net/ch/resource-service/pkg/models/ingress"
	"git.containerum.net/ch/resource-service/pkg/models/resources"
	"git.containerum.net/ch/resource-service/pkg/models/service"
	kubtypes "github.com/containerum/kube-client/pkg/model"
)

type DeployActions interface {
	GetDeploymentsList(ctx context.Context, nsID string) (deployment.DeploymentList, error)
	GetDeployment(ctx context.Context, nsID, deplName string) (*deployment.Deployment, error)
	CreateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) (*deployment.Deployment, error)
	UpdateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) (*deployment.Deployment, error)
	SetDeploymentReplicas(ctx context.Context, nsID, deplName string, req kubtypes.UpdateReplicas) (*deployment.Deployment, error)
	SetDeploymentContainerImage(ctx context.Context, nsID, deplName string, req kubtypes.UpdateImage) (*deployment.Deployment, error)
	DeleteDeployment(ctx context.Context, nsID, deplName string) error
	DeleteAllDeployments(ctx context.Context, nsID string) error
}

type DomainActions interface {
	GetDomainsList(ctx context.Context, page, per_page string) (domain.DomainList, error)
	GetDomain(ctx context.Context, domain string) (*domain.Domain, error)
	AddDomain(ctx context.Context, req domain.Domain) (*domain.Domain, error)
	DeleteDomain(ctx context.Context, domain string) error
}

type IngressActions interface {
	CreateIngress(ctx context.Context, nsID string, req kubtypes.Ingress) (*ingress.Ingress, error)
	GetIngressesList(ctx context.Context, nsID string) (ingress.IngressList, error)
	GetIngress(ctx context.Context, nsID, ingressName string) (*ingress.Ingress, error)
	UpdateIngress(ctx context.Context, nsID string, req kubtypes.Ingress) (*ingress.Ingress, error)
	DeleteIngress(ctx context.Context, nsID, ingressName string) error
	DeleteAllIngresses(ctx context.Context, nsID string) error
}

type ServiceActions interface {
	CreateService(ctx context.Context, nsID string, req kubtypes.Service) (*service.Service, error)
	GetServices(ctx context.Context, nsID string) (service.ServiceList, error)
	GetService(ctx context.Context, nsID, serviceName string) (*service.Service, error)
	UpdateService(ctx context.Context, nsID string, req kubtypes.Service) (*service.Service, error)
	DeleteService(ctx context.Context, nsID, serviceName string) error
	DeleteAllServices(ctx context.Context, nsID string) error
}

type ResourcesActions interface {
	GetResourcesCount(ctx context.Context) (*resources.GetResourcesCountResponse, error)
	DeleteAllResourcesInNamespace(ctx context.Context, nsID string) error
	DeleteAllUserResources(ctx context.Context) error
}
