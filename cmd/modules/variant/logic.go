package variant

import (
	"net/http"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
	md "app/cmd/models"
)

func (l *logic) List(filter md.VariantFilter) ([]md.FindVariant, common.Status, error) {
	results, status, err := Repository.List(filter)
	if err != nil {
		return []md.FindVariant{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(variant md.Variant) (md.Variant, common.Status, error) {
	result, status, err := Repository.Insert(variant)
	if err != nil {
		return md.Variant{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter md.VariantFilter, update md.VariantUpdate) (md.Variant, common.Status, error) {
	var variant md.Variant
	err := mapstructure.Decode(update, &variant)
	if err != nil {
		return md.Variant{}, http.StatusInternalServerError, err
	}

	result, status, err := Repository.Update(filter, variant)
	if err != nil {
		return md.Variant{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter md.VariantFilter) (md.VariantFilter, common.Status, error) {
	result, status, err := Repository.Archive(filter)
	if err != nil {
		return md.VariantFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter md.VariantFilter) (md.VariantFilter, common.Status, error) {
	result, status, err := Repository.Restore(filter)
	if err != nil {
		return md.VariantFilter{}, status, err
	}

	return result, status, nil
}
