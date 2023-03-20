package portal_user

import (
	"github.com/gofiber/fiber/v2"
)

type logic struct{}

var Logic logic

func (l *logic) List(filter PortalUserFilter) ([]PortalUser, []interface{}) {
	var errors []interface{} = nil

	//Validate PortalUser Struct
	validationError := filter.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
	}

	results, err := Repository.List(filter)
	if err != nil {
		errors = append(errors, err.Error())
	}

	return results, errors
}

func (l *logic) Insert(portal_user PortalUser) (PortalUser, []interface{}) {
	var errors []interface{} = nil

	//Validate PortalUser Struct
	validationError := portal_user.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
	}

	result, err := Repository.Insert(portal_user)
	if err != nil {
		errors = append(errors, err.Error())
	}

	return result, errors
}

func (l *logic) Update(filter PortalUserFilter, portal_user PortalUser) (PortalUser, []interface{}) {
	var errors []interface{} = nil

	updatePortalUser := PortalUserUpdate(portal_user)
	updateValidationError := updatePortalUser.Validate()
	if updateValidationError.Errors != nil {
		errors = append(errors, updateValidationError.Errors)
	}

	//Validate PortalUser Struct
	filterValidationError := filter.Validate()
	if filterValidationError.Errors != nil {
		errors = append(errors, filterValidationError.Errors)
	}

	result, err := Repository.Update(filter, portal_user)
	if err != nil {
		errors = append(errors, err.Error())
	}

	return result, errors
}

func (l *logic) Archive(c *fiber.Ctx) error {
	return Repository.Archive(c)
}
