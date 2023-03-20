package portal_user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/helper"
)

func GetHandler(c *fiber.Ctx) error {
	filter := new(PortalUserFilter)
	parseError := c.BodyParser(filter)
	if parseError != nil {
		return parseError
	}

	result, logicError := GetAllLogic(filter)
	if logicError != nil {
		return helper.JSON(c, logicError, true)
	}

	return helper.JSON(c, result)
}

func CreateHandler(c *fiber.Ctx) error {
	portal_user := new(PortalUser)
	parseError := c.BodyParser(portal_user)
	if parseError != nil {
		return parseError
	}

	result, logicError := CreateLogic(portal_user)
	if logicError != nil {
		return helper.JSON(c, logicError, true)
	}

	return helper.JSON(c, result)
}

func UpdateHandler(c *fiber.Ctx) error {
	return UpdateLogic(c)
}

func DeleteHandler(c *fiber.Ctx) error {
	return DeleteLogic(c)
}
