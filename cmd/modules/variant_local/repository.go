package variant_local

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	md "app/cmd/models"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter md.VariantLocalFilter, omits ...string) ([]md.FindVariantLocal, common.Status, error) {
	var variantLocals []md.FindVariantLocal

	result := pg.Conn.DB.Omit(omits...).Model(&md.VariantLocal{}).Preload("Model").Find(&variantLocals, filter)
	if result.Error != nil {
		return []md.FindVariantLocal{}, http.StatusInternalServerError, result.Error
	}

	return variantLocals, http.StatusOK, nil
}

func (r *repository) Insert(variantLocal md.VariantLocal) (md.VariantLocal, common.Status, error) {
	result := pg.Conn.DB.Model(&md.VariantLocal{}).Create(&variantLocal)
	if result.Error != nil {
		return md.VariantLocal{}, http.StatusInternalServerError, result.Error
	}

	return variantLocal, http.StatusOK, nil
}

func (r *repository) Update(filter md.VariantLocalFilter, variantLocal md.VariantLocal) (md.VariantLocal, common.Status, error) {
	result := pg.Conn.DB.Model(&md.VariantLocal{}).Where(filter).Updates(&variantLocal).Scan(&variantLocal)
	if result.Error != nil {
		return md.VariantLocal{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.VariantLocal{}, http.StatusNotFound, errors.New("can't find any variant local with this filter")
	}

	return variantLocal, http.StatusOK, nil
}

func (r *repository) Archive(filter md.VariantLocalFilter) (md.VariantLocalFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&md.VariantLocal{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.VariantLocalFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.VariantLocalFilter{}, http.StatusNotFound, errors.New("can't find any variant local with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter md.VariantLocalFilter) (md.VariantLocalFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&md.VariantLocal{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return md.VariantLocalFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return md.VariantLocalFilter{}, http.StatusNotFound, errors.New("can't find any variant local with this filter")
	}

	return filter, http.StatusOK, nil
}
