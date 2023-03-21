package portal_user

import (
	"net/http"
)

type logic struct{}

var Logic logic

func (l *logic) List(filter PortalUserFilter) ([]PortalUser, int, []interface{}) {
	var errors []interface{} = nil

	// Validate PortalUserFilter Struct
	validationError := filter.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
		return []PortalUser{}, http.StatusNotAcceptable, errors
	}

	results, status, err := Repository.List(filter)
	if err != nil {
		errors = append(errors, err.Error())
		return []PortalUser{}, status, errors
	}

	return results, status, errors
}

func (l *logic) Insert(portal_user PortalUser) (PortalUser, int, []interface{}) {
	var errors []interface{} = nil

	//Validate PortalUser Struct
	validationError := portal_user.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
		return PortalUser{}, http.StatusNotAcceptable, errors
	}

	result, status, err := Repository.Insert(portal_user)
	if err != nil {
		errors = append(errors, err.Error())
		return PortalUser{}, status, errors
	}

	return result, status, errors
}

func (l *logic) Update(filter PortalUserFilter, portal_user PortalUser) (PortalUser, int, []interface{}) {
	var errors []interface{} = nil

	updatePortalUser := PortalUserUpdate(portal_user)
	updateValidationError := updatePortalUser.Validate()
	if updateValidationError.Errors != nil {
		errors = append(errors, updateValidationError.Errors)
		return PortalUser{}, http.StatusNotAcceptable, errors
	}

	//Validate PortalUser Struct
	filterValidationError := filter.Validate()
	if filterValidationError.Errors != nil {
		errors = append(errors, filterValidationError.Errors)
		return PortalUser{}, http.StatusNotAcceptable, errors
	}

	result, status, err := Repository.Update(filter, portal_user)
	if err != nil {
		errors = append(errors, err.Error())
		return PortalUser{}, status, errors
	}

	return result, status, errors
}

func (l *logic) Archive(filter PortalUserFilter) (PortalUserFilter, int, []interface{}) {
	var errors []interface{} = nil

	// Validate PortalUserFilter Struct
	validationError := filter.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
		return PortalUserFilter{}, http.StatusNotAcceptable, errors
	}

	result, status, err := Repository.Archive(filter)
	if err != nil {
		errors = append(errors, err.Error())
		return PortalUserFilter{}, status, errors
	}

	return result, status, errors
}
