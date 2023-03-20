package portal_user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
)

func Routes(router fiber.Router) {
	pg.Conn.Migrate(&PortalUser{})

	userGroup := router.Group("/users")

	userGroup.Post("/find", GetHandler)
	userGroup.Post("/create", CreateHandler)
	userGroup.Patch("/update/:id", UpdateHandler)
	userGroup.Get("/archive/:id", DeleteHandler)
}
