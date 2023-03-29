package portal_user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
)

func Routes(router fiber.Router) {
	pg.Conn.Migrate(&PortalUser{})

	portalUserGroup := router.Group("/portal_user")

	portalUserGroup.Post("/find", Handler.List)
	portalUserGroup.Post("/create", Handler.Insert)
	portalUserGroup.Patch("/update", Handler.Update)
	portalUserGroup.Delete("/archive", Handler.Archive)
	portalUserGroup.Patch("/restore", Handler.Restore)
	portalUserGroup.Post("/login", Handler.Login)
	portalUserGroup.Post("/forgot", Handler.ForgotPassword)
	portalUserGroup.Post("/upload", Handler.Upload)
}
