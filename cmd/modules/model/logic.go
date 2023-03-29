package model

import (
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
)

func (l *logic) List(filter ModelFilter) ([]FindModel, common.Status, error) {
	if len(filter.Model) > 0 {
		filter.Model = strings.ToLower(filter.Model)
	}

	results, status, err := Repository.List(filter)
	if err != nil {
		return []FindModel{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(model Model) (Model, common.Status, error) {
	if len(model.Model) > 0 {
		model.Model = strings.ToLower(model.Model)
	}

	result, status, err := Repository.Insert(model)
	if err != nil {
		return Model{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter ModelFilter, update ModelUpdate) (Model, common.Status, error) {
	if len(filter.Model) > 0 {
		filter.Model = strings.ToLower(filter.Model)
	}

	if len(update.Model) > 0 {
		update.Model = strings.ToLower(update.Model)
	}

	var company Model
	err := mapstructure.Decode(update, &company)
	if err != nil {
		return Model{}, http.StatusInternalServerError, err
	}

	result, status, err := Repository.Update(filter, company)
	if err != nil {
		return Model{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter ModelFilter) (ModelFilter, common.Status, error) {
	if len(filter.Model) > 0 {
		filter.Model = strings.ToLower(filter.Model)
	}

	result, status, err := Repository.Archive(filter)
	if err != nil {
		return ModelFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter ModelFilter) (ModelFilter, common.Status, error) {
	if len(filter.Model) > 0 {
		filter.Model = strings.ToLower(filter.Model)
	}

	result, status, err := Repository.Restore(filter)
	if err != nil {
		return ModelFilter{}, status, err
	}

	return result, status, nil
}
