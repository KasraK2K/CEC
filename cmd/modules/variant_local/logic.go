package variant_local

import (
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
	md "app/cmd/models"
)

func (l *logic) List(filter md.VariantLocalFilter) ([]md.FindVariantLocal, common.Status, error) {
	if len(filter.Region) > 0 {
		filter.Region = strings.ToLower(filter.Region)
	}

	results, status, err := Repository.List(filter)
	if err != nil {
		return []md.FindVariantLocal{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(variantLocal md.VariantLocal) (md.VariantLocal, common.Status, error) {
	if len(variantLocal.Region) > 0 {
		variantLocal.Region = strings.ToLower(variantLocal.Region)
	}

	result, status, err := Repository.Insert(variantLocal)
	if err != nil {
		return md.VariantLocal{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter md.VariantLocalFilter, update md.VariantLocalUpdate) (md.VariantLocal, common.Status, error) {
	if len(filter.Region) > 0 {
		filter.Region = strings.ToLower(filter.Region)
	}

	if len(update.Region) > 0 {
		update.Region = strings.ToLower(update.Region)
	}

	var variantLocal md.VariantLocal
	err := mapstructure.Decode(update, &variantLocal)
	if err != nil {
		return md.VariantLocal{}, http.StatusInternalServerError, err
	}

	result, status, err := Repository.Update(filter, variantLocal)
	if err != nil {
		return md.VariantLocal{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter md.VariantLocalFilter) (md.VariantLocalFilter, common.Status, error) {
	if len(filter.Region) > 0 {
		filter.Region = strings.ToLower(filter.Region)
	}

	result, status, err := Repository.Archive(filter)
	if err != nil {
		return md.VariantLocalFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter md.VariantLocalFilter) (md.VariantLocalFilter, common.Status, error) {
	if len(filter.Region) > 0 {
		filter.Region = strings.ToLower(filter.Region)
	}

	result, status, err := Repository.Restore(filter)
	if err != nil {
		return md.VariantLocalFilter{}, status, err
	}

	return result, status, nil
}
