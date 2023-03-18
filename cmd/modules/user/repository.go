package user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
)

func GetAllUsersRepository(c *fiber.Ctx) error {
	return c.SendString("Get All Users")
}

func GetOneUserRepository(c *fiber.Ctx) error {
	return c.SendString("Get One User")
}

func CreateUserRepository(user *User) error {
	result := pg.Conn.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUserRepository(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

func DeleteUserRepository(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
