package variant_local

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
)

func Routes(router fiber.Router) {
	pg.Conn.Migrate(&VariantLocal{})

	variantLocalGroup := router.Group("/variant_local")

	variantLocalGroup.Post("/find", Handler.List)
	variantLocalGroup.Post("/create", Handler.Insert)
	variantLocalGroup.Patch("/update", Handler.Update)
	variantLocalGroup.Delete("/archive", Handler.Archive)
	variantLocalGroup.Patch("/restore", Handler.Restore)
}
