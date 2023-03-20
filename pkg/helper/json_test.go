package helper

import (
	"encoding/json"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"app/pkg/config"
)

type TestJsonStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMarshal(t *testing.T) {
	mockData := TestJsonStruct{
		Name: "Kasra",
		Age:  37,
	}

	t.Run("Development", func(t *testing.T) {
		config.AppConfig.MODE = "development"
		expectedData := mockData

		actualResult, err := Marshal(expectedData)
		if err != nil {
			t.Errorf("It has error on Marshal: %v", err)
		}
		expectedResult, _ := json.MarshalIndent(expectedData, "", "  ")

		if string(actualResult) != string(expectedResult) {
			t.Errorf("Actual result is not same as expected result")
		}
	})

	t.Run("Production", func(t *testing.T) {
		config.AppConfig.MODE = "production"
		expectedData := mockData

		actualResult, err := Marshal(expectedData)
		if err != nil {
			t.Errorf("It has error on Marshal: %v", err)
		}
		expectedResult, _ := json.Marshal(expectedData)

		if string(actualResult) != string(expectedResult) {
			t.Errorf("Actual result is not same as expected result")
		}
	})
}

func TestJSON(t *testing.T) {
	mockData := TestJsonStruct{
		Name: "Kasra",
		Age:  37,
	}

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	err := JSON(c, mockData)
	if err != nil {
		t.Errorf("error on JSON command: %v", err)
	}
}
