package company

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	md "app/cmd/models"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter md.CompanyFilter, omits ...string) ([]md.FindCompany, common.Status, error) {
	var companies []md.FindCompany

	result := pg.Conn.DB.Omit(omits...).Model(&md.Company{}).Preload("Model").Find(&companies, filter)
	if result.Error != nil {
		return []md.FindCompany{}, http.StatusInternalServerError, result.Error
	}

	return companies, http.StatusOK, nil
}

func (r *repository) Insert(company md.Company) (md.Company, common.Status, error) {
	result := pg.Conn.DB.Model(&md.Company{}).Create(&company)
	if result.Error != nil {
		return md.Company{}, http.StatusInternalServerError, result.Error
	}

	return company, http.StatusOK, nil
}

func (r *repository) Update(filter md.CompanyFilter, company md.Company) (md.Company, common.Status, error) {
	result := pg.Conn.DB.Model(&md.Company{}).Where(filter).Updates(&company).Scan(&company)
	if result.Error != nil {
		return md.Company{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.Company{}, http.StatusNotFound, errors.New("can't find any company with this filter")
	}

	return company, http.StatusOK, nil
}

func (r *repository) Archive(filter md.CompanyFilter) (md.CompanyFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&md.Company{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.CompanyFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.CompanyFilter{}, http.StatusNotFound, errors.New("can't find any company with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter md.CompanyFilter) (md.CompanyFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&md.Company{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.CompanyFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.CompanyFilter{}, http.StatusNotFound, errors.New("can't find any company with this filter")
	}

	return filter, http.StatusOK, nil
}
