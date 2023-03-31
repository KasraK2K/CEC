package portal_user

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	md "app/cmd/models"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter md.PortalUserFilter, omits ...string) ([]md.PortalUser, common.Status, error) {
	var portalUsers []md.PortalUser

	result := pg.Conn.DB.Omit(omits...).Model(&md.PortalUser{}).Find(&portalUsers, filter)
	if result.Error != nil {
		return []md.PortalUser{}, http.StatusInternalServerError, result.Error
	}

	return portalUsers, http.StatusOK, nil
}

func (r *repository) Insert(portalUser md.PortalUser) (md.PortalUser, common.Status, error) {
	result := pg.Conn.DB.Model(&md.PortalUser{}).Create(&portalUser)
	if result.Error != nil {
		return md.PortalUser{}, http.StatusInternalServerError, result.Error
	}

	portalUser.Password = ""
	return portalUser, http.StatusOK, nil
}

func (r *repository) Update(filter md.PortalUserFilter, update md.PortalUser) (md.PortalUser, common.Status, error) {
	result := pg.Conn.DB.Model(&md.PortalUser{}).Where(filter).Updates(&update).Scan(&update)
	if result.Error != nil {
		return md.PortalUser{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.PortalUser{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	update.Password = ""
	return update, http.StatusOK, nil
}

func (r *repository) Archive(filter md.PortalUserFilter) (md.PortalUserFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&md.PortalUser{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.PortalUserFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.PortalUserFilter{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter md.PortalUserFilter) (md.PortalUserFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&md.PortalUser{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.PortalUserFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.PortalUserFilter{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	return filter, http.StatusOK, nil
}
