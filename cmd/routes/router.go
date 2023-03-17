package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"app/cmd/modules/user"
)

func Routes(app *fiber.App) {
	app.Get("/_health", health)
	app.Get("/_metrics", monitor.New(monitor.Config{Title: "Default Metrics Page"}))

	v1 := app.Group("/v1")
	user.Routes(v1)
}

func health(c *fiber.Ctx) error {
	return c.SendString("Everything works fine")
}
