package main

import (
	"CEC/pkg/config"
	"CEC/pkg/storage/pg"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.SetConfig()

	app := fiber.New(fiber.Config{
		Prefork:       config.AppConfig.PREFORK,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Compare Electric Car v1.0.0",
	})

	// Call to fill DB in struct
	pg.Conn.Connect()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.AppConfig.PORT)))
}
