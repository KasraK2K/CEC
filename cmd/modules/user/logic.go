package user

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllUsersLogic(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All Users")
}

func GetOneUserLogic(ctx *fiber.Ctx) error {
	return ctx.SendString("Get One User")
}

func CreateUserLogic(user *User) (*User, []interface{}) {
	var errors []interface{} = nil

	//Validate User Struct
	validationError := user.Validate()
	if validationError.Errors != nil {
		errors = append(errors, validationError.Errors)
	}

	CreateUserRepository(user)
	return user, errors
}

func UpdateUserLogic(ctx *fiber.Ctx) error {
	return ctx.SendString("Update User")
}

func DeleteUserLogic(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete User")
}
