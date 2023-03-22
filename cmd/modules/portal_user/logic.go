package portal_user

import (
	"net/http"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
	"app/pkg/helper"
)

type logic struct{}

var Logic logic

func (l *logic) List(filter PortalUserFilter) ([]PortalUser, common.Status, []interface{}) {
	var errors []interface{} = nil

	// Validate PortalUserFilter Struct
	validationError := helper.Validator(filter)
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

func (l *logic) Insert(portalUser PortalUser) (PortalUser, common.Status, []interface{}) {
	var errors []interface{} = nil

	//Validate PortalUser Struct
	validationError := helper.Validator(portalUser)
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
		return PortalUser{}, http.StatusNotAcceptable, errors
	}

	// Hash password
	if len(portalUser.Password) > 0 {
		hash, err := helper.HashPassword(portalUser.Password)
		if err != nil {
			errors = append(errors, err.Error())
			return PortalUser{}, http.StatusInternalServerError, errors
		}
		portalUser.Password = hash
	}

	result, status, err := Repository.Insert(portalUser)
	if err != nil {
		errors = append(errors, err.Error())
		return PortalUser{}, status, errors
	}

	return result, status, errors
}

func (l *logic) Update(filter PortalUserFilter, portalUser PortalUser) (PortalUser, common.Status, []interface{}) {
	var errors []interface{} = nil

	var portalUserUpdate PortalUserUpdate
	err := mapstructure.Decode(portalUser, &portalUserUpdate)
	if err != nil {
		errors = append(errors, err.Error())
		return PortalUser{}, http.StatusInternalServerError, errors
	}

	updateValidationError := helper.Validator(portalUserUpdate)
	if updateValidationError.Errors != nil {
		errors = append(errors, updateValidationError.Errors)
		return PortalUser{}, http.StatusNotAcceptable, errors
	}

	//Validate PortalUser Struct
	filterValidationError := helper.Validator(filter)
	if filterValidationError.Errors != nil {
		errors = append(errors, filterValidationError.Errors)
		return PortalUser{}, http.StatusNotAcceptable, errors
	}

	// Hash password
	if len(portalUser.Password) > 0 {
		hash, err := helper.HashPassword(portalUser.Password)
		if err != nil {
			errors = append(errors, err.Error())
			return PortalUser{}, http.StatusInternalServerError, errors
		}
		portalUser.Password = hash
	}

	result, status, err := Repository.Update(filter, portalUser)
	if err != nil {
		errors = append(errors, err.Error())
		return PortalUser{}, status, errors
	}

	return result, status, errors
}

func (l *logic) Archive(filter PortalUserFilter) (PortalUserFilter, common.Status, []interface{}) {
	var errors []interface{} = nil

	// Validate PortalUserFilter Struct
	validationError := helper.Validator(filter)
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

func (l *logic) Restore(filter PortalUserFilter) (PortalUserFilter, common.Status, []interface{}) {
	var errors []interface{} = nil

	// Validate PortalUserFilter Struct
	validationError := helper.Validator(filter)
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

func (l *logic) Login(payload PortalUserLoginPayload) (string, common.Status, []interface{}) {
	var errors []interface{} = nil

	// Validate LoginPayload Struct
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
		return "", http.StatusNotAcceptable, errors
	}

	filter := PortalUserFilter{Email: payload.Email}
	results, status, err := Repository.List(filter)
	if err != nil {
		errors = append(errors, err.Error())
		return "", status, errors
	}

	if len(results) == 0 {
		errors = append(errors, "email or password is wrong")
		return "", http.StatusNotFound, errors
	}

	portalUser := results[0]
	if !helper.ComparePassword(portalUser.Password, payload.Password) {
		errors = append(errors, "email or password is wrong")
		return "", http.StatusNotFound, errors
	}

	payloadClaims := helper.PayloadClaims{
		ID:       portalUser.ID,
		RoleID:   1,
		Platform: uint8(payload.Platform),
		UserType: helper.Token.UserType.Portal,
	}
	token, err := helper.Token.CreateToken(payloadClaims)
	if err != nil {
		errors = append(errors, err.Error())
		return "", http.StatusInternalServerError, errors
	}

	return token, http.StatusOK, nil
}
