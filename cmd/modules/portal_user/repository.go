package portal_user

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter PortalUserFilter, omits ...string) ([]PortalUser, common.Status, error) {
	var portalUsers []PortalUser

	result := pg.Conn.DB.Omit(omits...).Model(&PortalUser{}).Find(&portalUsers, filter)
	if result.Error != nil {
		return []PortalUser{}, http.StatusInternalServerError, result.Error
	}

	return portalUsers, http.StatusOK, nil
}

func (r *repository) Insert(portalUser PortalUser) (PortalUser, common.Status, error) {
	result := pg.Conn.DB.Model(&PortalUser{}).Create(&portalUser)
	if result.Error != nil {
		return PortalUser{}, http.StatusInternalServerError, result.Error
	}

	portalUser.Password = ""
	return portalUser, http.StatusOK, nil
}

func (r *repository) Update(filter PortalUserFilter, update PortalUserUpdate) (PortalUserUpdate, common.Status, error) {
	result := pg.Conn.DB.Model(&PortalUser{}).Where(filter).Updates(&update).Scan(&update)
	if result.Error != nil {
		return PortalUserUpdate{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return PortalUserUpdate{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	update.Password = ""
	return update, http.StatusOK, nil
}

func (r *repository) Archive(filter PortalUserFilter) (PortalUserFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&PortalUser{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return PortalUserFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return PortalUserFilter{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter PortalUserFilter) (PortalUserFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&PortalUser{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return PortalUserFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return PortalUserFilter{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	return filter, http.StatusOK, nil
}
