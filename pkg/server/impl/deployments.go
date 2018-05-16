package impl

import (
	"context"

	"git.containerum.net/ch/resource-service/pkg/db"
	"git.containerum.net/ch/resource-service/pkg/models/deployment"
	"git.containerum.net/ch/resource-service/pkg/rsErrors"
	"git.containerum.net/ch/resource-service/pkg/server"
	"github.com/containerum/cherry/adaptors/cherrylog"
	kubtypes "github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/utils/httputil"
	"github.com/sirupsen/logrus"
)

type DeployActionsImpl struct {
	mongo *db.MongoStorage
	log   *cherrylog.LogrusAdapter
}

func NewDeployActionsImpl(mongo *db.MongoStorage) *DeployActionsImpl {
	return &DeployActionsImpl{
		mongo: mongo,
		log:   cherrylog.NewLogrusAdapter(logrus.WithField("component", "deploy_actions")),
	}
}

func (da *DeployActionsImpl) GetDeploymentsList(ctx context.Context, nsID string) ([]deployment.Deployment, error) {
	userID := httputil.MustGetUserID(ctx)
	da.log.WithFields(logrus.Fields{
		"user_id":   userID,
		"namespace": nsID,
	}).Info("get deployments")

	return da.mongo.GetDeploymentList(nsID)
}

func (da *DeployActionsImpl) GetDeployment(ctx context.Context, nsID, deplName string) (*deployment.Deployment, error) {
	userID := httputil.MustGetUserID(ctx)
	da.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"ns_id":       nsID,
		"deploy_name": deplName,
	}).Info("get deployment by label")

	ret, err := da.mongo.GetDeploymentByName(nsID, deplName)

	return &ret, err
}

func (da *DeployActionsImpl) CreateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) (*deployment.Deployment, error) {
	userID := httputil.MustGetUserID(ctx)
	da.log.WithFields(logrus.Fields{
		"user_id": userID,
		"ns_id":   nsID,
	}).Info("create deployment")

	//TODO
	/*
		nsUsage, getErr := da.NamespaceDB(tx).GetNamespaceUsage(ctx, ns.Namespace)
		if getErr != nil {
			return getErr
		}

		if chkErr := server.CheckDeploymentCreateQuotas(ns.Namespace, nsUsage, deploy); chkErr != nil {
			return chkErr
		}*/

	if err := server.CalculateDeployResources(&deploy); err != nil {
		return nil, err
	}

	createdDeploy, err := da.mongo.CreateDeployment(deployment.DeploymentFromKube(nsID, userID, deploy))
	if err != nil {
		return nil, err
	}

	//TODO
	/*	if err = da.Kube.DeleteDeployment(ctx, nsID, deplName); delErr != nil {
		return err
	}*/

	return &createdDeploy, nil
}

func (da *DeployActionsImpl) DeleteDeployment(ctx context.Context, nsID, deplName string) error {
	userID := httputil.MustGetUserID(ctx)
	da.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"ns_id":       nsID,
		"deploy_name": deplName,
	}).Info("delete deployment")

	err := da.mongo.DeleteDeployment(nsID, deplName)
	if err != nil {
		return err
	}

	//TODO
	/*if delErr = da.Kube.DeleteDeployment(ctx, nsID, deplName); delErr != nil {
		return delErr
	}*/

	return nil
}

func (da *DeployActionsImpl) UpdateDeployment(ctx context.Context, nsID string, deploy kubtypes.Deployment) (*deployment.Deployment, error) {
	userID := httputil.MustGetUserID(ctx)
	da.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"ns_id":       nsID,
		"deploy_name": deploy.Name,
	}).Infof("replacing deployment with %#v", deploy)

	if err := server.CalculateDeployResources(&deploy); err != nil {
		return nil, err
	}

	//TODO
	/*
		nsUsage, getErr := da.NamespaceDB(tx).GetNamespaceUsage(ctx, ns.Namespace)
		if getErr != nil {
			return getErr
		}

		if chkErr := server.CheckDeploymentReplaceQuotas(ns.Namespace, nsUsage, oldDeploy, deploy); chkErr != nil {
			return chkErr
		}

		if replaceErr := da.Kube.ReplaceDeployment(ctx, ns.ID, deploy); replaceErr != nil {
			return replaceErr
		}
	*/

	err := da.mongo.UpdateDeployment(deployment.DeploymentFromKube(nsID, userID, deploy))
	if err != nil {
		return nil, err
	}

	updatedDeploy, err := da.mongo.GetDeploymentByName(nsID, deploy.Name)
	if err != nil {
		return nil, err
	}

	return &updatedDeploy, nil
}

func (da *DeployActionsImpl) SetDeploymentReplicas(ctx context.Context, nsID, deplName string, req kubtypes.UpdateReplicas) (*deployment.Deployment, error) {
	userID := httputil.MustGetUserID(ctx)
	da.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"ns_id":       nsID,
		"deploy_name": deplName,
	}).Infof("set deployment replicas %#v", req)

	//TODO
	/*
		nsUsage, getErr := da.NamespaceDB(tx).GetNamespaceUsage(ctx, ns.Namespace)
		if getErr != nil {
			return getErr
		}

		if chkErr := server.CheckDeploymentReplicasChangeQuotas(ns.Namespace, nsUsage, deploy, req.Replicas); chkErr != nil {
			return chkErr
		}

		if setErr := da.DeployDB(tx).SetDeploymentReplicas(ctx, userID, nsLabel, deplName, req.Replicas); setErr != nil {
			return setErr
		}

		if setErr := da.Kube.SetDeploymentReplicas(ctx, ns.ID, deplName, req.Replicas); setErr != nil {
			return setErr
		}
	*/

	oldDeploy, err := da.mongo.GetDeploymentByName(nsID, deplName)
	if err != nil {
		return nil, err
	}

	oldDeploy.Replicas = req.Replicas

	err = da.mongo.UpdateDeployment(oldDeploy)
	if err != nil {
		return nil, err
	}

	updatedDeploy, err := da.mongo.GetDeploymentByName(nsID, deplName)
	if err != nil {
		return nil, err
	}

	return &updatedDeploy, nil
}

func (da *DeployActionsImpl) SetDeploymentContainerImage(ctx context.Context, nsID, deplName string, req kubtypes.UpdateImage) (*deployment.Deployment, error) {
	userID := httputil.MustGetUserID(ctx)
	da.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"ns_id":       nsID,
		"deploy_name": deplName,
	}).Infof("set container image %#v", req)

	//TODO
	/*
		if setErr := da.DeployDB(tx).SetContainerImage(ctx, userID, nsLabel, deplName, req); setErr != nil {
			return setErr
		}

		setErr := da.Kube.SetContainerImage(ctx, nsID, deplName, req)
		if setErr != nil {
			return setErr
		}
	*/

	oldDeploy, err := da.mongo.GetDeploymentByName(nsID, deplName)
	if err != nil {
		return nil, err
	}

	updated := false
	for i, c := range oldDeploy.Containers {
		if c.Name == req.Container {
			oldDeploy.Containers[i].Image = req.Image
			updated = true
			break
		}

	}
	if !updated {
		return nil, rserrors.ErrInternal().AddDetails("No image found")
	}

	err = da.mongo.UpdateDeployment(oldDeploy)
	if err != nil {
		return nil, err
	}

	updatedDeploy, err := da.mongo.GetDeploymentByName(nsID, deplName)
	if err != nil {
		return nil, err
	}

	return &updatedDeploy, nil
}
