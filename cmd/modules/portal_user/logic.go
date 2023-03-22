package portal_user

import (
	"net/http"

	"app/pkg/helper"
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

	results, status, err := Repository.List(filter, []string{"password"}...)
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

	// Hash password
	if len(portal_user.Password) > 0 {
		hash, err := helper.HashPassword(portal_user.Password)
		if err != nil {
			errors = append(errors, err.Error())
			return PortalUser{}, http.StatusInternalServerError, errors
		}
		portal_user.Password = hash
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

	// Hash password
	if len(portal_user.Password) > 0 {
		hash, err := helper.HashPassword(portal_user.Password)
		if err != nil {
			errors = append(errors, err.Error())
			return PortalUser{}, http.StatusInternalServerError, errors
		}
		portal_user.Password = hash
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

func (l *logic) Restore(filter PortalUserFilter) (PortalUserFilter, int, []interface{}) {
	var errors []interface{} = nil

	// Validate PortalUserFilter Struct
	validationError := filter.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
		return PortalUserFilter{}, http.StatusNotAcceptable, errors
	}

	result, status, err := Repository.Restore(filter)
	if err != nil {
		errors = append(errors, err.Error())
		return PortalUserFilter{}, status, errors
	}

	return result, status, errors
}

func (l *logic) Login(email, password string) (string, int, []interface{}) {
	var errors []interface{} = nil

	filter := PortalUserFilter{Email: email}
	results, status, err := Repository.List(filter)
	if err != nil {
		errors = append(errors, err.Error())
		return "", status, errors
	}

	if len(results) == 0 {
		errors = append(errors, "email or password is wrong")
		return "", http.StatusNotFound, errors
	}

	portal_user := results[0]
	if !helper.ComparePassword(portal_user.Password, password) {
		errors = append(errors, "email or password is wrong")
		return "", http.StatusNotFound, errors
	}

	payloadClaims := helper.PayloadClaims{
		ID:     portal_user.ID,
		RoleID: 1,
	}
	token, err := helper.CreateToken(payloadClaims)
	if err != nil {
		errors = append(errors, err.Error())
		return "", http.StatusInternalServerError, errors
	}

	return token, http.StatusOK, nil
}
