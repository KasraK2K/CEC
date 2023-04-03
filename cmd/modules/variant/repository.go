package variant

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	md "app/cmd/models"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter md.VariantFilter, omits ...string) ([]md.FindVariant, common.Status, error) {
	var variants []md.FindVariant

	result := pg.Conn.DB.Omit(omits...).Model(&md.Variant{}).Preload("Model").Preload("Company").Find(&variants, filter)
	if result.Error != nil {
		return []md.FindVariant{}, http.StatusInternalServerError, result.Error
	}

	return variants, http.StatusOK, nil
}

func (r *repository) Insert(variant md.Variant) (md.Variant, common.Status, error) {
	result := pg.Conn.DB.Model(&md.Variant{}).Create(&variant)
	if result.Error != nil {
		return md.Variant{}, http.StatusInternalServerError, result.Error
	}

	return variant, http.StatusOK, nil
}

func (r *repository) Update(filter md.VariantFilter, variant md.Variant) (md.Variant, common.Status, error) {
	result := pg.Conn.DB.Model(&md.Variant{}).Where(filter).Updates(&variant).Scan(&variant)
	if result.Error != nil {
		return md.Variant{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.Variant{}, http.StatusNotFound, errors.New("can't find any variant with this filter")
	}

	return variant, http.StatusOK, nil
}

func (r *repository) Archive(filter md.VariantFilter) (md.VariantFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&md.Variant{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.VariantFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.VariantFilter{}, http.StatusNotFound, errors.New("can't find any variant with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter md.VariantFilter) (md.VariantFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&md.Variant{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.VariantFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.VariantFilter{}, http.StatusNotFound, errors.New("can't find any variant with this filter")
	}

	return filter, http.StatusOK, nil
}
