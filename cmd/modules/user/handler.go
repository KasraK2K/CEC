package user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/helper"
)

func GetAllUsersHandler(ctx *fiber.Ctx) error {
	return GetAllUsersLogic(ctx)
}

func GetOneUserHandler(ctx *fiber.Ctx) error {
	return GetOneUserLogic(ctx)
}

func CreateUserHandler(ctx *fiber.Ctx) error {
	user := new(User)
	parseError := ctx.BodyParser(user)
	if parseError != nil {
		return parseError
	}

	result, logicError := CreateUserLogic(user)
	if logicError != nil {
		return helper.JSON(ctx, logicError, true)
	}

	return helper.JSON(ctx, result)
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	return UpdateUserLogic(ctx)
}

func DeleteUserHandler(ctx *fiber.Ctx) error {
	return DeleteUserLogic(ctx)
}
