package company

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter CompanyFilter, omits ...string) ([]Company, common.Status, error) {
	var portalUsers []Company

	result := pg.Conn.DB.Omit(omits...).Model(&Company{}).Find(&portalUsers, filter)
	if result.Error != nil {
		return []Company{}, http.StatusInternalServerError, result.Error
	}

	return portalUsers, http.StatusOK, nil
}

func (r *repository) Insert(company Company) (Company, common.Status, error) {
	result := pg.Conn.DB.Model(&Company{}).Create(&company)
	if result.Error != nil {
		return Company{}, http.StatusInternalServerError, result.Error
	}

	return company, http.StatusOK, nil
}

func (r *repository) Update(filter CompanyFilter, company Company) (Company, common.Status, error) {
	result := pg.Conn.DB.Model(&Company{}).Where(filter).Updates(&company).Scan(&company)
	if result.Error != nil {
		return Company{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return Company{}, http.StatusNotFound, errors.New("can't find any company with this filter")
	}

	return company, http.StatusOK, nil
}

func (r *repository) Archive(filter CompanyFilter) (CompanyFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&Company{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return CompanyFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return CompanyFilter{}, http.StatusNotFound, errors.New("can't find any company with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter CompanyFilter) (CompanyFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&Company{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return CompanyFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return CompanyFilter{}, http.StatusNotFound, errors.New("can't find any company with this filter")
	}

	return filter, http.StatusOK, nil
}
