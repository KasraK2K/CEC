package middleware

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/valyala/fasthttp"

	"app/pkg/config"
	"app/pkg/helper"
)

func TestPullOutToken(t *testing.T) {
	config.SetConfig()

	app := fiber.New()
	app.Use(PullOutToken)

	claimsPayload := helper.PayloadClaims{
		ID:         1,
		Permission: "111",
		Platform:   1,
		UserType:   1,
	}
	token, err := helper.Token.CreateToken(claimsPayload)
	if err != nil {
		t.Errorf("TestPullOutToken error on sign token")
	}

	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	app.Handler()(ctx)

	tokenPayload := ctx.UserValue("TokenPayload")

	var payload helper.Payload
	mapstructure.Decode(tokenPayload, &payload)

	if payload.Permission != claimsPayload.Permission || payload.ID != claimsPayload.ID || payload.Platform != claimsPayload.Platform || payload.UserType != claimsPayload.UserType {
		t.Errorf("TestPullOutToken error on finding TestPullOutToken")
	}
}
