package postgres

import (
	"context"

	"database/sql"

	"git.containerum.net/ch/json-types/misc"
	rstypes "git.containerum.net/ch/json-types/resource-service"
	"git.containerum.net/ch/kube-client/pkg/cherry/resource-service"
	"git.containerum.net/ch/resource-service/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func (db *pgDB) isVolumeExists(ctx context.Context, userID, label string) (exists bool, err error) {
	params := map[string]interface{}{
		"user_id": userID,
		"label":   label,
	}
	entry := db.log.WithFields(params)
	entry.Debug("check if volume exists")

	var count int
	query, args, _ := sqlx.Named( /* language=sql */
		`SELECT count(v.*)
		FROM volumes v
		JOIN permissions p ON p.resource_id = v.id AND p.kind = 'volume'
		WHERE p.user_id = :user_id AND p.resource_label = :label`, params)
	err = sqlx.GetContext(ctx, db.extLog, &count, db.extLog.Rebind(query), args...)
	if err != nil {
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}

	entry.Debugf("found %d volumes", count)
	exists = count > 0
	return
}

func (db *pgDB) addVolumesToNamespaces(ctx context.Context,
	nsIDs []string, nsMap map[string]rstypes.NamespaceWithVolumes) (err error) {
	db.log.Debugf("add volumes to namespaces %v", nsIDs)
	type volWithNsID struct {
		rstypes.VolumeWithPermission
		NsID string `db:"ns_id"`
	}
	volsWithNsID := make([]volWithNsID, 0)
	query, args, _ := sqlx.In( /* language=sql */
		`SELECT v.*, 
			p.id AS perm_id,
			p.kind,
			p.resource_id,
			p.resource_label,
			p.owner_user_id,
			p.create_time,
			p.user_id,
			p.access_level,
			p.limited,
			p.access_level_change_time,
			p.new_access_level,
			d.ns_id
		FROM volumes v
		JOIN volume_mounts vm ON v.id = vm.volume_id
		JOIN containers c ON vm.container_id = c.id
		JOIN deployments d ON c.depl_id = d.id
		JOIN permissions p ON p.resource_id = v.id
		WHERE d.ns_id IN (?)`, nsIDs)
	err = sqlx.SelectContext(ctx, db.extLog, &volsWithNsID, db.extLog.Rebind(query), args...)
	switch err {
	case nil, sql.ErrNoRows:
		err = nil
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}

	// fetch non-persistent volumes
	query, args, _ = sqlx.In( /* language=sql */
		`SELECT v.*, 
			p.id AS perm_id,
			p.kind,
			p.resource_id,
			p.resource_label,
			p.owner_user_id,
			p.create_time,
			p.user_id,
			p.access_level,
			p.limited,
			p.access_level_change_time,
			p.new_access_level,
			v.ns_id
		FROM volumes v
		JOIN permissions p ON p.resource_id = v.id
		WHERE v.ns_id IN (?)`, nsIDs)
	npvs := make([]volWithNsID, 0)
	err = sqlx.SelectContext(ctx, db.extLog, &volsWithNsID, db.extLog.Rebind(query), args...)
	switch err {
	case nil, sql.ErrNoRows:
		err = nil
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}

	volsWithNsID = append(volsWithNsID, npvs...)

	for _, v := range volsWithNsID {
		ns := nsMap[v.NsID]
		ns.Volume = append(ns.Volume, v.VolumeWithPermission)
		nsMap[v.NsID] = ns
	}

	return
}

func (db *pgDB) CreateVolume(ctx context.Context, userID, label string, volume *rstypes.Volume) (err error) {
	db.log.WithFields(logrus.Fields{
		"user_id": userID,
		"label":   label,
	}).Infof("creating volume %#v", volume)

	var exists bool
	if exists, err = db.isVolumeExists(ctx, userID, label); err != nil {
		return
	}
	if exists {
		err = rserrors.ErrResourceAlreadyExists().Log(err, db.log)
		return
	}

	query, args, _ := sqlx.Named( /* language=sql */
		`INSERT INTO volumes
		(
			tariff_id,
			capacity,
			replicas,
			ns_id
		)
		VALUES (:tariff_id, :capacity, :replicas, :ns_id)
		RETURNING *`,
		volume)
	err = sqlx.GetContext(ctx, db.extLog, volume, db.extLog.Rebind(query), args...)
	if err != nil {
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}

	_, err = sqlx.NamedExecContext(ctx, db.extLog, /* language=sql */
		`INSERT INTO permissions
		(
			kind,
			resource_id,
			resource_label,
			owner_user_id,
			user_id
		)
		VALUES ('volume', :resource_id, :resource_label, :user_id, :user_id)`,
		rstypes.PermissionRecord{
			ResourceID:    misc.WrapString(volume.ID),
			ResourceLabel: label,
			UserID:        userID,
		})
	if err != nil {
		err = rserrors.ErrDatabase().Log(err, db.log)
	}

	return
}

func (db *pgDB) GetUserVolumes(ctx context.Context,
	userID string, filters *models.VolumeFilterParams) (ret []rstypes.VolumeWithPermission, err error) {
	ret = make([]rstypes.VolumeWithPermission, 0)
	db.log.WithField("user_id", userID).Debugf("get user volumes (filters %#v)", filters)

	params := struct {
		UserID string `db:"user_id"`
		*models.VolumeFilterParams
	}{
		UserID:             userID,
		VolumeFilterParams: filters,
	}
	query, args, _ := sqlx.Named( /* language=sql */
		`SELECT v.*,
			p.id AS perm_id,
			p.kind,
			p.resource_id,
			p.resource_label,
			p.owner_user_id,
			p.create_time,
			p.user_id,
			p.access_level,
			p.limited,
			p.access_level_change_time,
			p.new_access_level
		FROM volumes v
		JOIN permissions p ON p.resource_id = v.id AND p.kind = 'volume'
		WHERE 
			p.user_id = :user_id AND
			(NOT v.deleted OR NOT :not_deleted) AND
			(v.deleted OR NOT :deleted) AND
			(p.limited OR NOT :limited) AND
			(NOT p.limited OR NOT :not_limited) AND
			(p.owner_user_id = p.user_id OR NOT :owned) AND
			(v.ns_id IS NULL OR NOT :persistent) AND
			(v.ns_id IS NOT NULL OR NOT :not_persistent)
		ORDER BY v.create_time DESC`,
		params)

	err = sqlx.SelectContext(ctx, db.extLog, &ret, db.extLog.Rebind(query), args...)
	switch err {
	case nil, sql.ErrNoRows:
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
	}

	return
}

func (db *pgDB) GetAllVolumes(ctx context.Context,
	page, perPage int, filters *models.VolumeFilterParams) (ret []rstypes.VolumeWithPermission, err error) {
	ret = make([]rstypes.VolumeWithPermission, 0)

	db.log.WithFields(logrus.Fields{
		"page":     page,
		"per_page": perPage,
	}).Debug("get all volumes")

	params := struct {
		Limit  int `db:"limit"`
		Offset int `db:"offset"`
		*models.VolumeFilterParams
	}{
		Limit:              perPage,
		Offset:             (page - 1) * perPage,
		VolumeFilterParams: filters,
	}
	query, args, _ := sqlx.Named( /* language=sql */
		`SELECT v.*, 
			p.id AS perm_id,
			p.kind,
			p.resource_id,
			p.resource_label,
			p.owner_user_id,
			p.create_time,
			p.user_id,
			p.access_level,
			p.limited,
			p.access_level_change_time,
			p.new_access_level
			FROM volumes v
			JOIN permissions p ON p.resource_id = v.id AND p.kind = 'volume'
			WHERE 
				(NOT v.deleted OR NOT :not_deleted) AND
				(v.deleted OR NOT :deleted) AND
				(p.limited OR NOT :limited) AND
				(NOT p.limited OR NOT :not_limited) AND
				(p.owner_user_id = p.user_id OR NOT :owned) AND
				(v.ns_id IS NULL OR NOT :persistent) AND
				(v.ns_id IS NOT NULL OR NOT :not_persistent)
			ORDER BY v.create_time DESC
			LIMIT :limit
			OFFSET :offset`,
		params)

	err = sqlx.SelectContext(ctx, db.extLog, &ret, db.extLog.Rebind(query), args...)
	switch err {
	case nil, sql.ErrNoRows:
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
	}

	return
}

func (db *pgDB) GetUserVolumeByLabel(ctx context.Context,
	userID, label string) (ret rstypes.VolumeWithPermission, err error) {
	params := map[string]interface{}{
		"user_id": userID,
		"label":   label,
	}
	db.log.WithFields(params).Debug("get user volume by label")

	query, args, _ := sqlx.Named( /* language=sql */
		`SELECT v.*, 
			p.id AS perm_id,
			p.kind,
			p.resource_id,
			p.resource_label,
			p.owner_user_id,
			p.create_time,
			p.user_id,
			p.access_level,
			p.limited,
			p.access_level_change_time,
			p.new_access_level
		FROM volumes v
		JOIN permissions p ON p.resource_id = v.id AND p.kind = 'volume'
		WHERE p.user_id = :user_id AND p.resource_label = :label`,
		params)
	err = sqlx.GetContext(ctx, db.extLog, &ret, db.extLog.Rebind(query), args...)
	switch err {
	case nil:
	case sql.ErrNoRows:
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
	}

	return
}

func (db *pgDB) GetVolumeWithUserPermissions(ctx context.Context,
	userID, label string) (ret rstypes.VolumeWithUserPermissions, err error) {
	params := map[string]interface{}{
		"user_id": userID,
		"label":   label,
	}
	db.log.WithFields(params).Debug("get volume with user permissions")

	ret.Users = make([]rstypes.PermissionRecord, 0)

	query, args, _ := sqlx.Named( /* language=sql */
		`SELECT v.*,
			p.id AS perm_id,
			p.kind,
			p.resource_id,
			p.resource_label,
			p.owner_user_id,
			p.create_time,
			p.user_id,
			p.access_level,
			p.limited,
			p.access_level_change_time,
			p.new_access_level
		FROM volumes v
		JOIN permissions p ON p.resource_id = v.id AND p.kind = 'volume'
		WHERE p.user_id = :user_id AND p.resource_label = :label`,
		params)
	err = sqlx.GetContext(ctx, db.extLog, &ret.VolumeWithPermission, db.extLog.Rebind(query), args...)
	switch err {
	case nil:
	case sql.ErrNoRows:
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
		return
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}

	query, args, _ = sqlx.Named( /* language=sql */
		`SELECT
			p.id AS perm_id,
			p.kind,
			p.resource_id,
			p.resource_label,
			p.owner_user_id,
			p.create_time,
			p.user_id,
			p.access_level,
			p.limited,
			p.access_level_change_time,
			p.new_access_level
		FROM permissions p
		WHERE owner_user_id != user_id AND
				kind = 'volume' AND
				resource_id = :id`,
		ret.Resource)
	err = sqlx.SelectContext(ctx, db.extLog, &ret.Users, db.extLog.Rebind(query), args...)
	switch err {
	case nil:
	case sql.ErrNoRows:
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
	}

	return
}

func (db *pgDB) GetVolumesLinkedWithUserNamespace(ctx context.Context, userID, label string) (ret []rstypes.VolumeWithPermission, err error) {
	params := map[string]interface{}{
		"user_id":        userID,
		"resource_label": label,
	}
	db.log.WithFields(params).Debug("get volumes linked with user namespace")

	nsID, err := db.getNamespaceID(ctx, userID, label)
	if err != nil {
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}
	if nsID == "" {
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
		return
	}

	ret = make([]rstypes.VolumeWithPermission, 0)

	query, args, _ := sqlx.Named( /* language=sql */
		`SELECT v.*,
			p.id AS perm_id,
			p.kind,
			p.resource_id,
			p.resource_label,
			p.owner_user_id,
			p.create_time,
			p.user_id,
			p.access_level,
			p.limited,
			p.access_level_change_time,
			p.new_access_level
		FROM volumes v
		JOIN volume_mounts vm ON v.id = vm.volume_id
		JOIN containers c ON vm.container_id = c.id
		JOIN deployments d ON c.depl_id = d.id
		JOIN permissions p ON p.resource_id = d.ns_id AND p.kind = 'namespace'
		WHERE p.user_id = :user_id AND p.resource_label = :label`,
		params)
	err = sqlx.SelectContext(ctx, db.extLog, &ret, db.extLog.Rebind(query), args...)
	switch err {
	case nil, sql.ErrNoRows:
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
	}

	return
}

func (db *pgDB) DeleteUserVolumeByLabel(ctx context.Context, userID, label string) (volume rstypes.Volume, err error) {
	params := map[string]interface{}{
		"user_id":        userID,
		"resource_label": label,
	}
	db.log.WithFields(params).Debug("delete user volume by label")

	query, args, _ := sqlx.Named( /* language=sql */
		`WITH user_vol AS (
			SELECT resource_id
			FROM permissions
			WHERE user_id = owner_user_id AND
					kind = 'volume' AND
					user_id = :user_id AND 
					resource_label = :resource_label
		)
		UPDATE volumes
		SET deleted = TRUE, active = FALSE 
		WHERE id IN (SELECT resource_id FROM user_vol)
		RETURNING *`,
		params)
	err = sqlx.GetContext(ctx, db.extLog, &volume, db.extLog.Rebind(query), args...)
	switch err {
	case nil:
	case sql.ErrNoRows:
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
		return
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}

	return
}

func (db *pgDB) DeleteAllUserVolumes(ctx context.Context, userID string, nonPersistentOnly bool) (ret []rstypes.Volume, err error) {
	params := map[string]interface{}{
		"user_id":             userID,
		"non_persistent_only": nonPersistentOnly,
	}
	db.log.WithFields(params).Debug("delete all user volumes")

	query, args, _ := sqlx.Named( /* language=sql */
		`WITH user_vol AS (
			SELECT resource_id
			FROM permissions
			WHERE user_id = owner_user_id AND 
					kind = 'volume' AND 
					user_id = :user_id
		)
		UPDATE volumes
		SET deleted = TRUE, active = FALSE
		WHERE id IN (SELECT resource_id FROM user_vol) AND (ns_id IS NOT NULL OR NOT :non_persistent_only)
		RETURNING *`,
		params)
	ret = make([]rstypes.Volume, 0)
	err = sqlx.SelectContext(ctx, db.extLog, &ret, db.extLog.Rebind(query), args...)
	switch err {
	case nil, sql.ErrNoRows:
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}

	return
}

func (db *pgDB) RenameVolume(ctx context.Context, userID, oldLabel, newLabel string) (err error) {
	params := map[string]interface{}{
		"user_id":   userID,
		"old_label": oldLabel,
		"new_label": newLabel,
	}
	db.log.WithFields(params).Debug("rename user volume")

	exists, err := db.isVolumeExists(ctx, userID, newLabel)
	if err != nil {
		return
	}
	if exists {
		err = rserrors.ErrResourceAlreadyExists().Log(err, db.log)
		return
	}

	result, err := sqlx.NamedExecContext(ctx, db.extLog, /* language=sql */
		`UPDATE permissions
		SET resource_label = :old_label
		WHERE owner_user_id = :user_id AND
				kind = 'volume' AND
				resource_label = :new_label`,
		params)
	if err != nil {
		err = rserrors.ErrDatabase().Log(err, db.log)
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
	}

	return
}

func (db *pgDB) ResizeVolume(ctx context.Context, volume *rstypes.Volume) (err error) {
	db.log.WithField("volume_id", volume.ID).Debugf("update volume to %#v", volume)

	query, args, _ := sqlx.Named( /* language=sql */
		`UPDATE volumes
		SET
			tariff_id = :tariff_id,
			capacity = :capacity,
			replicas = :replicas
		WHERE id = :id`,
		volume)
	err = sqlx.GetContext(ctx, db.extLog, volume, db.extLog.Rebind(query), args...)
	switch err {
	case nil:
	case sql.ErrNoRows:
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
		return
	default:
		err = rserrors.ErrDatabase().Log(err, db.log)
		return
	}

	return
}

func (db *pgDB) SetVolumeActiveByID(ctx context.Context, id string, active bool) (err error) {
	params := map[string]interface{}{
		"id":     id,
		"active": active,
	}
	db.log.WithFields(params).Debug("activating volume by id")

	result, err := sqlx.NamedExecContext(ctx, db.extLog, /* language=sql */
		`UPDATE volumes SET active = :id WHERE id = :active`, params)
	if err != nil {
		err = rserrors.ErrDatabase().Log(err, db.log)
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
	}

	return
}

func (db *pgDB) SetUserVolumeActive(ctx context.Context, userID, label string, active bool) (err error) {
	params := map[string]interface{}{
		"user_id": userID,
		"label":   label,
		"active":  active,
	}
	db.log.WithFields(params).Debug("activating user volume")

	result, err := sqlx.NamedExecContext(ctx, db.extLog, /* language=sql */
		`WITH user_vol AS (
			SELECT resource_id
			FROM permissions
			WHERE owner_user_id = user_id AND 
				user_id = $1 AND 
				kind = 'volume' AND
				resource_label = $2
		)
		UPDATE volumes 
		SET active = $2 
		WHERE id IN (SELECT resource_id FROM user_vol)`,
		params)
	if err != nil {
		err = rserrors.ErrDatabase().Log(err, db.log)
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		err = rserrors.ErrResourceNotExists().Log(err, db.log)
	}

	return
}
