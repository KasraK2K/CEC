package portal_user

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"app/pkg/storage/pg"
)

type repository struct{}

var Repository repository

func (r *repository) List(filter PortalUserFilter) ([]PortalUser, error) {
	var portalUsers []PortalUser
	result := pg.Conn.DB.Find(&portalUsers, filter)
	if result.Error != nil {
		return []PortalUser{}, result.Error
	}
	return portalUsers, nil
}

func (r *repository) Insert(portal_user PortalUser) (PortalUser, error) {
	result := pg.Conn.DB.Create(&portal_user)
	if result.Error != nil {
		return PortalUser{}, result.Error
	}
	return portal_user, nil
}

func (r *repository) Update(filter interface{}, portal_user PortalUser) (PortalUser, error) {
	result := pg.Conn.DB.Model(&PortalUser{}).Where(filter).Updates(&portal_user).Scan(&portal_user)
	if result.Error != nil {
		return PortalUser{}, result.Error
	}

	if result.RowsAffected == 0 {
		return PortalUser{}, errors.New("can't find any user with this filter")
	}

	return portal_user, nil
}

func (r *repository) Archive(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
