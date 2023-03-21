package portal_user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
)

func Routes(router fiber.Router) {
	pg.Conn.Migrate(&PortalUser{})

	userGroup := router.Group("/users")

	userGroup.Post("/find", Handler.List)
	userGroup.Post("/create", Handler.Insert)
	userGroup.Patch("/update", Handler.Update)
	userGroup.Delete("/archive", Handler.Archive)
}
