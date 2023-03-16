package user

import (
	"github.com/gofiber/fiber/v2"
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
		return ctx.JSON(logicError)
	}

	return ctx.JSON(result)
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	return UpdateUserLogic(ctx)
}

func DeleteUserHandler(ctx *fiber.Ctx) error {
	return DeleteUserLogic(ctx)
}
