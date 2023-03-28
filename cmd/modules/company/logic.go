package company

import (
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
)

func (l *logic) List(filter CompanyFilter) ([]Company, common.Status, error) {
	if len(filter.Brand) > 0 {
		filter.Brand = strings.ToLower(filter.Brand)
	}

	results, status, err := Repository.List(filter, []string{"password"}...)
	if err != nil {
		return []Company{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(company Company) (Company, common.Status, error) {
	result, status, err := Repository.Insert(company)
	if err != nil {
		return Company{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter CompanyFilter, update CompanyUpdate) (Company, common.Status, error) {
	if len(filter.Brand) > 0 {
		filter.Brand = strings.ToLower(filter.Brand)
	}

	var company Company
	err := mapstructure.Decode(update, &company)
	if err != nil {
		return Company{}, http.StatusInternalServerError, err
	}

	result, status, err := Repository.Update(filter, company)
	if err != nil {
		return Company{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter CompanyFilter) (CompanyFilter, common.Status, error) {
	if len(filter.Brand) > 0 {
		filter.Brand = strings.ToLower(filter.Brand)
	}

	result, status, err := Repository.Archive(filter)
	if err != nil {
		return CompanyFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter CompanyFilter) (CompanyFilter, common.Status, error) {
	if len(filter.Brand) > 0 {
		filter.Brand = strings.ToLower(filter.Brand)
	}

	result, status, err := Repository.Restore(filter)
	if err != nil {
		return CompanyFilter{}, status, err
	}

	return result, status, nil
}
