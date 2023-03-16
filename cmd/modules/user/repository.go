package user

import (
	"CEC/pkg/storage/pg"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsersRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All Users")
}

func GetOneUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Get One User")
}

func CreateUserRepository(user *User) {
	pg.Conn.DB.Create(&user)
}

func UpdateUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Update User")
}

func DeleteUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete User")
}
