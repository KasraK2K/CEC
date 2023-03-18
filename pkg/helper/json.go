package helper

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"

	"app/pkg/config"
)

func Marshal(v any) ([]byte, error) {
	if config.AppConfig.MODE == "development" {
		return json.MarshalIndent(v, "", "  ")
	} else {
		return json.Marshal(v)
	}
}

func JSON(ctx *fiber.Ctx, data any, errors ...bool) error {
	responseData := AddMetaData(data, errors...)
	byteData, err := Marshal(responseData)
	if err != nil {
		log.Panic(err)
	}

	return ctx.SendString(string(byteData))
}
