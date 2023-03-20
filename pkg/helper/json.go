package helper

import (
	"encoding/json"
	"log"

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

func JSON(c *fiber.Ctx, data any, errors ...bool) error {
	metadata := AddMetaData(data, errors...)
	byteData, err := Marshal(metadata)
	if err != nil {
		log.Panic(err)
	}

	return c.SendString(string(byteData))
}
