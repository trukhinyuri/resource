package db

import (
	"time"

	"git.containerum.net/ch/resource-service/pkg/models/configmap"
	"git.containerum.net/ch/resource-service/pkg/rsErrors"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

func (mongo *MongoStorage) GetConfigMap(namespaceID, cmName string) (configmap.ConfigMapResource, error) {
	mongo.logger.Debugf("getting configmap")
	var collection = mongo.db.C(CollectionCM)
	var result configmap.ConfigMapResource
	var err error
	if err = collection.Find(configmap.OneSelectQuery(namespaceID, cmName)).One(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get configmap")
		if err == mgo.ErrNotFound {
			return result, rserrors.ErrResourceNotExists().AddDetails(cmName)
		}
		return result, PipErr{err}.ToMongerr().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetConfigMapList(namespaceID string) (configmap.ConfigMapList, error) {
	mongo.logger.Debugf("getting configmaps list")
	var collection = mongo.db.C(CollectionCM)
	result := make(configmap.ConfigMapList, 0)
	if err := collection.Find(bson.M{
		"namespaceid": namespaceID,
		"deleted":     false,
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get configmaps list")
		return result, PipErr{err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

// If ID is empty, then generates UUID4 and uses it
func (mongo *MongoStorage) CreateConfigMap(cm configmap.ConfigMapResource) (configmap.ConfigMapResource, error) {
	mongo.logger.Debugf("creating configmap")
	var collection = mongo.db.C(CollectionCM)
	if cm.ID == "" {
		cm.ID = uuid.New().String()
	}
	cm.Data = nil
	cm.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	if err := collection.Insert(cm); err != nil {
		mongo.logger.WithError(err).Errorf("unable to create configmap")
		if mgo.IsDup(err) {
			return cm, rserrors.ErrResourceAlreadyExists().AddDetailsErr(err)
		}
		return cm, PipErr{err}.ToMongerr().Extract()
	}
	return cm, nil
}

func (mongo *MongoStorage) DeleteConfigMap(namespaceID, name string) error {
	mongo.logger.Debugf("deleting configmap")
	var collection = mongo.db.C(CollectionCM)
	err := collection.Update(configmap.ConfigMapResource{
		ConfigMap: model.ConfigMap{
			Name: name,
		},
		NamespaceID: namespaceID,
	}.OneSelectQuery(),
		bson.M{
			"$set": bson.M{"deleted": true,
				"configmap.deletedat": time.Now().UTC().Format(time.RFC3339)},
		})
	if err != nil {
		mongo.logger.WithError(err).Errorf("unable to delete configmap")
		if err == mgo.ErrNotFound {
			return rserrors.ErrResourceNotExists().AddDetails(name)
		}
		return PipErr{error: err}.ToMongerr().Extract()
	}
	return nil
}

func (mongo *MongoStorage) DeleteAllConfigMapsInNamespace(namespaceID string) error {
	mongo.logger.Debugf("deleting all configmaps in namespace")
	var collection = mongo.db.C(CollectionCM)
	_, err := collection.UpdateAll(configmap.ConfigMapResource{
		NamespaceID: namespaceID,
	}.AllSelectQuery(),
		bson.M{
			"$set": bson.M{"deleted": true,
				"configmap.deletedat": time.Now().UTC().Format(time.RFC3339)},
		})
	if err != nil {
		mongo.logger.WithError(err).Errorf("unable to delete configmap")
		return PipErr{error: err}.ToMongerr().Extract()
	}
	return nil
}

func (mongo *MongoStorage) DeleteAllConfigMapsByOwner(owner string) error {
	mongo.logger.Debugf("deleting all configmaps in namespace")
	var collection = mongo.db.C(CollectionCM)
	_, err := collection.UpdateAll(configmap.ConfigMapResource{
		ConfigMap: model.ConfigMap{Owner: owner},
	}.AllSelectOwnerQuery(),
		bson.M{
			"$set": bson.M{"deleted": true,
				"configmap.deletedat": time.Now().UTC().Format(time.RFC3339)},
		})
	if err != nil {
		mongo.logger.WithError(err).Errorf("unable to delete configmaps")
		return PipErr{error: err}.ToMongerr().Extract()
	}
	return nil
}

func (mongo *MongoStorage) RestoreConfigMap(namespaceID, name string) error {
	mongo.logger.Debugf("restoring configmap")
	var collection = mongo.db.C(CollectionCM)
	err := collection.Update(configmap.ConfigMapResource{
		ConfigMap: model.ConfigMap{
			Name: name,
		},
		NamespaceID: namespaceID,
	}.OneSelectDeletedQuery(),
		bson.M{
			"$set": bson.M{"deleted": false,
				"configmap.deletedat": ""},
		})
	if err != nil {
		mongo.logger.WithError(err).Errorf("unable to restore configmap")
		if err == mgo.ErrNotFound {
			return rserrors.ErrResourceNotExists().AddDetails(name)
		}
		return PipErr{error: err}.ToMongerr().Extract()
	}
	return nil
}

func (mongo *MongoStorage) CountConfigMaps(owner string) (int, error) {
	mongo.logger.Debugf("counting configmaps")
	var collection = mongo.db.C(CollectionCM)
	n, err := collection.Find(bson.M{
		"configmap.owner": owner,
		"deleted":         false,
	}).Count()
	if err != nil {
		return 0, PipErr{err}.ToMongerr().NotFoundToNil().Extract()
	}
	return n, nil
}

func (mongo *MongoStorage) CountAllConfigMaps() (int, error) {
	mongo.logger.Debugf("counting all configmaps")
	var collection = mongo.db.C(CollectionCM)
	n, err := collection.Find(bson.M{
		"deleted": false,
	}).Count()
	if err != nil {
		return 0, PipErr{err}.ToMongerr().NotFoundToNil().Extract()
	}
	return n, nil
}
