package user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
)

func GetAllUsersRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All Users")
}

func GetOneUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Get One User")
}

func CreateUserRepository(user *User) error {
	result := pg.Conn.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Update User")
}

func DeleteUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete User")
}
