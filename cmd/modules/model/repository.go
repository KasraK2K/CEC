package model

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	md "app/cmd/models"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter md.ModelFilter, omits ...string) ([]md.FindModel, common.Status, error) {
	var models []md.FindModel

	// var myData []map[string]interface{}

	result := pg.Conn.DB.
		Omit(omits...).
		Model(&md.Model{}).
		Preload("Company").
		Preload("VariantLocal").
		Find(&models, filter)
	if result.Error != nil {
		return []md.FindModel{}, http.StatusInternalServerError, result.Error
	}

	return models, http.StatusOK, nil
}

func (r *repository) Insert(model md.Model) (md.Model, common.Status, error) {
	result := pg.Conn.DB.Model(&md.Model{}).Create(&model)
	if result.Error != nil {
		return md.Model{}, http.StatusInternalServerError, result.Error
	}

	return model, http.StatusOK, nil
}

func (r *repository) Update(filter md.ModelFilter, model md.Model) (md.Model, common.Status, error) {
	result := pg.Conn.DB.Model(&md.Model{}).Where(filter).Updates(&model).Scan(&model)
	if result.Error != nil {
		return md.Model{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.Model{}, http.StatusNotFound, errors.New("can't find any model with this filter")
	}

	return model, http.StatusOK, nil
}

func (r *repository) Archive(filter md.ModelFilter) (md.ModelFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&md.Model{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.ModelFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.ModelFilter{}, http.StatusNotFound, errors.New("can't find any model with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter md.ModelFilter) (md.ModelFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&md.Model{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.ModelFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.ModelFilter{}, http.StatusNotFound, errors.New("can't find any model with this filter")
	}

	return filter, http.StatusOK, nil
}
