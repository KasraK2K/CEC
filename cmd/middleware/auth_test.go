package middleware

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/valyala/fasthttp"

	"app/pkg/config"
	"app/pkg/helper"
)

func TestPullOutToken(t *testing.T) {
	config.SetConfig("../../pkg/config/.env")

	app := fiber.New()
	app.Use(PullOutToken)

	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicGVybWlzc2lvbiI6IjExMTExMTExMTExIiwicGxhdGZvcm0iOjEsInVzZXJfdHlwZSI6MSwic3ViIjoiMSIsImV4cCI6MTY4MDg3Njc4MCwiaWF0IjoxNjgwNzkwMzgwfQ.6H5OIpJdSse0SrXGZe0sO2_FURFVD6BkU_GcTtuJewQ")

	app.Handler()(ctx)

	tokenPayload := ctx.UserValue("TokenPayload")

	var payload helper.Payload
	mapstructure.Decode(tokenPayload, &payload)

	if payload.Permission == "" || payload.ID == 0 || payload.Platform == 0 || payload.UserType == 0 {
		t.Errorf("TestPullOutToken error on finding TestPullOutToken")
	}
}
