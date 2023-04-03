package variant

import (
	"github.com/gofiber/fiber/v2"

	md "app/cmd/models"
	"app/pkg/storage/pg"
)

func Routes(router fiber.Router) {
	pg.Conn.Migrate(&md.Variant{})

	variantGroup := router.Group("/variant")

	variantGroup.Post("/find", Handler.List)
	variantGroup.Post("/create", Handler.Insert)
	variantGroup.Patch("/update", Handler.Update)
	variantGroup.Delete("/archive", Handler.Archive)
	variantGroup.Patch("/restore", Handler.Restore)
}
