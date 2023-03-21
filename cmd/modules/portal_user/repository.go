package portal_user

import (
	"errors"
	"net/http"

	"gorm.io/gorm/clause"

	"app/pkg/storage/pg"
)

type repository struct{}

var Repository repository

func (r *repository) List(filter PortalUserFilter) ([]PortalUser, int, error) {
	var portalUsers []PortalUser
	result := pg.Conn.DB.Find(&portalUsers, filter)
	if result.Error != nil {
		return []PortalUser{}, http.StatusInternalServerError, result.Error
	}
	return portalUsers, http.StatusOK, nil
}

func (r *repository) Insert(portal_user PortalUser) (PortalUser, int, error) {
	result := pg.Conn.DB.Create(&portal_user)
	if result.Error != nil {
		return PortalUser{}, http.StatusInternalServerError, result.Error
	}
	return portal_user, http.StatusOK, nil
}

func (r *repository) Update(filter interface{}, portal_user PortalUser) (PortalUser, int, error) {
	result := pg.Conn.DB.Model(&PortalUser{}).Where(filter).Updates(&portal_user).Scan(&portal_user)
	if result.Error != nil {
		return PortalUser{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return PortalUser{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	return portal_user, http.StatusOK, nil
}

func (r *repository) Archive(filter interface{}) (PortalUser, int, error) {
	var portal_user PortalUser
	result := pg.Conn.DB.Clauses(clause.Returning{}).Order("created_at desc").Take(&portal_user, filter).Delete(&portal_user)

	if result.Error != nil {
		return portal_user, http.StatusInternalServerError, result.Error
	}

	return portal_user, http.StatusOK, nil
}
