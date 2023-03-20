package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"app/cmd/routes"
	"app/pkg/config"
	"app/pkg/storage/pg"
)

func main() {
	config.SetConfig()
	pg.Conn.Connect()

	app := fiber.New(fiber.Config{
		Prefork:       config.AppConfig.PREFORK,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Compare Electric Car v1.0.0",
	})

	// Fiber Middleware
	app.Use(cache.New())
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{Max: 100, Expiration: 60 * time.Second}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Router
	routes.Routes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.AppConfig.PORT)))
}
