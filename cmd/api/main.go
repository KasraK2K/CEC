package main

import (
	"CEC/cmd/routes"
	"CEC/pkg/config"
	"CEC/pkg/storage/pg"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.SetConfig()
	pg.Conn.Connect()

	app := fiber.New(fiber.Config{
		Prefork:       config.AppConfig.PREFORK,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Compare Electric Car v1.0.0",
	})

	// Middleware

	// Router
	routes.Routes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.AppConfig.PORT)))
}
