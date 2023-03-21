package portal_user

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/pkg/storage/pg"
)

type repository struct{}

var Repository repository

func (r *repository) List(filter PortalUserFilter) ([]PortalUser, int, error) {
	var portalUsers []PortalUser
	result := pg.Conn.DB.Model(&PortalUser{}).Find(&portalUsers, filter)
	if result.Error != nil {
		return []PortalUser{}, http.StatusInternalServerError, result.Error
	}
	return portalUsers, http.StatusOK, nil
}

func (r *repository) Insert(portal_user PortalUser) (PortalUser, int, error) {
	result := pg.Conn.DB.Model(&PortalUser{}).Create(&portal_user)
	if result.Error != nil {
		return PortalUser{}, http.StatusInternalServerError, result.Error
	}
	return portal_user, http.StatusOK, nil
}

func (r *repository) Update(filter PortalUserFilter, portal_user PortalUser) (PortalUser, int, error) {
	result := pg.Conn.DB.Model(&PortalUser{}).Where(filter).Updates(&portal_user).Scan(&portal_user)
	if result.Error != nil {
		return PortalUser{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return PortalUser{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	return portal_user, http.StatusOK, nil
}

func (r *repository) Archive(filter PortalUserFilter) (PortalUserFilter, int, error) {
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

func (r *repository) Restore(filter PortalUserFilter) (PortalUserFilter, int, error) {
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
