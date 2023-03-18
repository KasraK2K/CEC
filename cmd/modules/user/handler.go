package user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/helper"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
	return GetAllUsersLogic(c)
}

func GetOneUserHandler(c *fiber.Ctx) error {
	return GetOneUserLogic(c)
}

func CreateUserHandler(c *fiber.Ctx) error {
	user := new(User)
	parseError := c.BodyParser(user)
	if parseError != nil {
		return parseError
	}

	result, logicError := CreateUserLogic(user)
	if logicError != nil {
		return helper.JSON(c, logicError, true)
	}

	return helper.JSON(c, result)
}

func UpdateUserHandler(c *fiber.Ctx) error {
	return UpdateUserLogic(c)
}

func DeleteUserHandler(c *fiber.Ctx) error {
	return DeleteUserLogic(c)
}
