package portal_user

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
)

func GetAllRepository(filter *PortalUserFilter) ([]PortalUser, error) {
	var portalUsers []PortalUser
	result := pg.Conn.DB.Find(&portalUsers, filter)
	if result.Error != nil {
		return []PortalUser{}, result.Error
	}
	return portalUsers, nil
}

func CreateRepository(portal_user *PortalUser) error {
	result := pg.Conn.DB.Create(&portal_user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRepository(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

func DeleteRepository(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
