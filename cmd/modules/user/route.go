package user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
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
