package model

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter ModelFilter, omits ...string) ([]FindModel, common.Status, error) {
	var models []FindModel

	result := pg.Conn.DB.
		Model(&Model{}).
		Omit(omits...).
		Preload("Company").
		Joins("LEFT JOIN companies ON companies.company_id = models.company_id").
		// Preload("VariantLocal").
		// Joins("JOIN variant_locals ON variant_locals.model_id = models.model_id").
		Find(&models, filter)

	// result := pg.Conn.DB.
	// 	Table("companies c").
	// 	Select("c.*, m.*, vl.*").
	// 	Joins("JOIN models m ON c.ID = m.CompanyID").
	// 	Joins("JOIN variant_locals vl ON m.ID = vl.ModelID").
	// 	Where("m.IsArchive = false AND vl.IsArchive = false").
	// 	Find(&models)
	if result.Error != nil {
		return []FindModel{}, http.StatusInternalServerError, result.Error
	}

	return models, http.StatusOK, nil
}

func (r *repository) Insert(model Model) (Model, common.Status, error) {
	result := pg.Conn.DB.Model(&Model{}).Create(&model)
	if result.Error != nil {
		return Model{}, http.StatusInternalServerError, result.Error
	}

	return model, http.StatusOK, nil
}

func (r *repository) Update(filter ModelFilter, model Model) (Model, common.Status, error) {
	result := pg.Conn.DB.Model(&Model{}).Where(filter).Updates(&model).Scan(&model)
	if result.Error != nil {
		return Model{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return Model{}, http.StatusNotFound, errors.New("can't find any model with this filter")
	}

	return model, http.StatusOK, nil
}

func (r *repository) Archive(filter ModelFilter) (ModelFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&Model{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return ModelFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return ModelFilter{}, http.StatusNotFound, errors.New("can't find any model with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter ModelFilter) (ModelFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&Model{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return ModelFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return ModelFilter{}, http.StatusNotFound, errors.New("can't find any model with this filter")
	}

	return filter, http.StatusOK, nil
}
