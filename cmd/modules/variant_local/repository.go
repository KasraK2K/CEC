package variant_local

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/cmd/common"
	"app/pkg/storage/pg"
)

func (r *repository) List(filter VariantLocalFilter, omits ...string) ([]VariantLocal, common.Status, error) {
	var variantLocals []VariantLocal

	result := pg.Conn.DB.Omit(omits...).Model(&VariantLocal{}).Find(&variantLocals, filter)
	if result.Error != nil {
		return []VariantLocal{}, http.StatusInternalServerError, result.Error
	}

	return variantLocals, http.StatusOK, nil
}

func (r *repository) Insert(variantLocal VariantLocal) (VariantLocal, common.Status, error) {
	result := pg.Conn.DB.Model(&VariantLocal{}).Create(&variantLocal)
	if result.Error != nil {
		return VariantLocal{}, http.StatusInternalServerError, result.Error
	}

	return variantLocal, http.StatusOK, nil
}

func (r *repository) Update(filter VariantLocalFilter, variantLocal VariantLocal) (VariantLocal, common.Status, error) {
	result := pg.Conn.DB.Model(&VariantLocal{}).Where(filter).Updates(&variantLocal).Scan(&variantLocal)
	if result.Error != nil {
		return VariantLocal{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return VariantLocal{}, http.StatusNotFound, errors.New("can't find any variant local with this filter")
	}

	return variantLocal, http.StatusOK, nil
}

func (r *repository) Archive(filter VariantLocalFilter) (VariantLocalFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Conn.DB.Model(&VariantLocal{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return VariantLocalFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return VariantLocalFilter{}, http.StatusNotFound, errors.New("can't find any variant local with this filter")
	}

	return filter, http.StatusOK, nil
}

func (r *repository) Restore(filter VariantLocalFilter) (VariantLocalFilter, common.Status, error) {
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Conn.DB.Unscoped().Model(&VariantLocal{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return VariantLocalFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return VariantLocalFilter{}, http.StatusNotFound, errors.New("can't find any variant local with this filter")
	}

	return filter, http.StatusOK, nil
}
