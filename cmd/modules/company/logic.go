package company

import (
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
	md "app/cmd/models"
)

func (l *logic) List(filter md.CompanyFilter) ([]md.FindCompany, common.Status, error) {
	if len(filter.Brand) > 0 {
		filter.Brand = strings.ToLower(filter.Brand)
	}

	results, status, err := Repository.List(filter)
	if err != nil {
		return []md.FindCompany{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(company md.Company) (md.Company, common.Status, error) {
	if len(company.Brand) > 0 {
		company.Brand = strings.ToLower(company.Brand)
	}

	result, status, err := Repository.Insert(company)
	if err != nil {
		return md.Company{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter md.CompanyFilter, update md.CompanyUpdate) (md.Company, common.Status, error) {
	if len(filter.Brand) > 0 {
		filter.Brand = strings.ToLower(filter.Brand)
	}

	if len(update.Brand) > 0 {
		update.Brand = strings.ToLower(update.Brand)
	}

	var company md.Company
	err := mapstructure.Decode(update, &company)
	if err != nil {
		return md.Company{}, http.StatusInternalServerError, err
	}

	result, status, err := Repository.Update(filter, company)
	if err != nil {
		return md.Company{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter md.CompanyFilter) (md.CompanyFilter, common.Status, error) {
	if len(filter.Brand) > 0 {
		filter.Brand = strings.ToLower(filter.Brand)
	}

	result, status, err := Repository.Archive(filter)
	if err != nil {
		return md.CompanyFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter md.CompanyFilter) (md.CompanyFilter, common.Status, error) {
	if len(filter.Brand) > 0 {
		filter.Brand = strings.ToLower(filter.Brand)
	}

	result, status, err := Repository.Restore(filter)
	if err != nil {
		return md.CompanyFilter{}, status, err
	}

	return result, status, nil
}
