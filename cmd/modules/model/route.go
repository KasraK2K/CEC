package model

import (
	"github.com/gofiber/fiber/v2"

	"app/cmd/middleware"
	"app/pkg/storage/pg"
)

func Routes(router fiber.Router) {
	pg.Conn.Migrate(&Model{})

	modelGroup := router.Group("/model")

	// Needed Permissions
	modelGroup.Use([]string{"/find", "/create"}, func(c *fiber.Ctx) error {
		return middleware.CheckPermission(c, 2)
	})

	// Routes
	modelGroup.Post("/find", Handler.List)
	modelGroup.Post("/create", Handler.Insert)
	modelGroup.Patch("/update", Handler.Update)
	modelGroup.Delete("/archive", Handler.Archive)
	modelGroup.Patch("/restore", Handler.Restore)
}
