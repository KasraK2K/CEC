package user

import (
	"CEC/pkg/storage/pg"

	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	pg.Conn.Migrate(&User{})

	userGroup := router.Group("/users")

	userGroup.Get("/", GetAllUsersHandler)
	userGroup.Get("/:id", GetOneUserHandler)
	userGroup.Post("/", CreateUserHandler)
	userGroup.Patch("/:id", UpdateUserHandler)
	userGroup.Get("/:id", DeleteUserHandler)
}
