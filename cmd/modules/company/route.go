package company

import (
	"github.com/gofiber/fiber/v2"

	md "app/cmd/models"
	"app/pkg/storage/pg"
)

func Routes(router fiber.Router) {
	pg.Conn.Migrate(&md.Company{})

	companyGroup := router.Group("/company")

	companyGroup.Post("/find", Handler.List)
	companyGroup.Post("/create", Handler.Insert)
	companyGroup.Patch("/update", Handler.Update)
	companyGroup.Delete("/archive", Handler.Archive)
	companyGroup.Patch("/restore", Handler.Restore)
}
