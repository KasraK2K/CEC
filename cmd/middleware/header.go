package middleware

import (
	"github.com/gofiber/fiber/v2"

	"app/pkg/config"
)

func AddMiddleware(c *fiber.Ctx) error {
	c.Response().Header.SetCanonical([]byte("Backend-Version"), []byte(config.AppConfig.APP_VERSION))
	c.Response().Header.SetCanonical([]byte("Frontend-Version"), []byte(config.AppConfig.APP_VERSION))
	c.Response().Header.SetCanonical([]byte("App-Version"), []byte(config.AppConfig.APP_VERSION))
	c.Response().Header.SetCanonical([]byte("Mode"), []byte(config.AppConfig.APP_VERSION))
	return c.Next()
}
