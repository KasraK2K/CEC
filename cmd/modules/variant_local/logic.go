package variant_local

import (
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
)

func (l *logic) List(filter VariantLocalFilter) ([]VariantLocal, common.Status, error) {
	if len(filter.Region) > 0 {
		filter.Region = strings.ToLower(filter.Region)
	}

	results, status, err := Repository.List(filter)
	if err != nil {
		return []VariantLocal{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(variantLocal VariantLocal) (VariantLocal, common.Status, error) {
	if len(variantLocal.Region) > 0 {
		variantLocal.Region = strings.ToLower(variantLocal.Region)
	}

	result, status, err := Repository.Insert(variantLocal)
	if err != nil {
		return VariantLocal{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter VariantLocalFilter, update VariantLocalUpdate) (VariantLocal, common.Status, error) {
	if len(filter.Region) > 0 {
		filter.Region = strings.ToLower(filter.Region)
	}

	if len(update.Region) > 0 {
		update.Region = strings.ToLower(update.Region)
	}

	var variantLocal VariantLocal
	err := mapstructure.Decode(update, &variantLocal)
	if err != nil {
		return VariantLocal{}, http.StatusInternalServerError, err
	}

	result, status, err := Repository.Update(filter, variantLocal)
	if err != nil {
		return VariantLocal{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter VariantLocalFilter) (VariantLocalFilter, common.Status, error) {
	if len(filter.Region) > 0 {
		filter.Region = strings.ToLower(filter.Region)
	}

	result, status, err := Repository.Archive(filter)
	if err != nil {
		return VariantLocalFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter VariantLocalFilter) (VariantLocalFilter, common.Status, error) {
	if len(filter.Region) > 0 {
		filter.Region = strings.ToLower(filter.Region)
	}

	result, status, err := Repository.Restore(filter)
	if err != nil {
		return VariantLocalFilter{}, status, err
	}

	return result, status, nil
}
