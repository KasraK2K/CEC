package portal_user

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllLogic(filter *PortalUserFilter) ([]PortalUser, []interface{}) {
	var errors []interface{} = nil

	//Validate PortalUser Struct
	validationError := filter.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
	}

	results, err := GetAllRepository(*filter)
	if err != nil {
		errors = append(errors, err.Error())
	}

	return results, errors
}

func CreateLogic(portal_user *PortalUser) (PortalUser, []interface{}) {
	var errors []interface{} = nil

	//Validate PortalUser Struct
	validationError := portal_user.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
	}

	result, err := CreateRepository(*portal_user)
	if err != nil {
		errors = append(errors, err.Error())
	}

	return result, errors
}

func UpdateLogic(c *fiber.Ctx) error {
	return UpdateRepository(c)
}

func DeleteLogic(c *fiber.Ctx) error {
	return DeleteRepository(c)
}
