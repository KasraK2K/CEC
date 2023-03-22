package helper

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"app/cmd/common"
	"app/pkg/config"
)

func Marshal(v interface{}) ([]byte, error) {
	if config.AppConfig.MODE == "development" {
		return json.MarshalIndent(v, "", "  ")
	} else {
		return json.Marshal(v)
	}
}

func JSON(c *fiber.Ctx, data interface{}, status common.Status) error {
	metadata := AddMetaData(data, int(status))
	byteData, err := Marshal(metadata)
	if err != nil {
		log.Panic(err)
	}

	c.Status(int(status))
	return c.SendString(string(byteData))
}
