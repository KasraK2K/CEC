package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"app/cmd/modules/portal_user"
	"app/pkg/helper"
)

func Routes(app *fiber.App) {
	app.Get("/_health", health)
	app.Get("/_metrics", monitor.New(monitor.Config{Title: "Default Metrics Page"}))

	v1 := app.Group("/v1")
	portal_user.Routes(v1)

	// Handle other routes
	app.Use("*", func(c *fiber.Ctx) error {
		return helper.JSON(c, "this route is not exist", http.StatusNotFound)
	})
}

func health(c *fiber.Ctx) error {
	return c.SendString("Everything works fine")
}
