package helper

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"app/pkg/config"
)

func Marshal(v interface{}) ([]byte, error) {
	if config.AppConfig.MODE == "development" {
		return json.MarshalIndent(v, "", "  ")
	} else {
		return json.Marshal(v)
	}
}

func JSON(c *fiber.Ctx, data interface{}, status int) error {
	metadata := AddMetaData(data, status)
	byteData, err := Marshal(metadata)
	if err != nil {
		log.Panic(err)
	}

	c.Status(status)
	return c.SendString(string(byteData))
}
