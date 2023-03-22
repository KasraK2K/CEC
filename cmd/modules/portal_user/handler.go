package portal_user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"app/pkg/helper"
)

type handler struct{}

var Handler handler

func (h *handler) List(c *fiber.Ctx) error {
	type JsonData struct {
		Filter PortalUserFilter `json:"filter"`
	}
	var payload JsonData
	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return helper.JSON(c, parseError.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	results, status, err := Logic.List(payload.Filter)
	if err != nil {
		return helper.JSON(c, err, status)
	}

	return helper.JSON(c, results, status)
}

func (h *handler) Insert(c *fiber.Ctx) error {
	type jsonData struct {
		Data PortalUser `json:"data"`
	}
	var payload jsonData

	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return helper.JSON(c, parseError.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	portalUser := payload.Data
	result, status, err := Logic.Insert(portalUser)
	if err != nil {
		return helper.JSON(c, err, status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Update(c *fiber.Ctx) error {
	type JsonData struct {
		Filter PortalUserFilter `json:"filter"`
		Data   PortalUser       `json:"data"`
	}
	var payload JsonData
	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return helper.JSON(c, parseError.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	portalUser := payload.Data
	result, status, err := Logic.Update(filter, portalUser)
	if err != nil {
		return helper.JSON(c, err, status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Archive(c *fiber.Ctx) error {
	type JsonData struct {
		Filter PortalUserFilter `json:"filter"`
	}
	var payload JsonData
	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return helper.JSON(c, parseError.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	result, status, err := Logic.Archive(filter)
	if err != nil {
		return helper.JSON(c, err, status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Restore(c *fiber.Ctx) error {
	type JsonData struct {
		Filter PortalUserFilter `json:"filter"`
	}
	var payload JsonData
	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return helper.JSON(c, parseError.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	result, status, err := Logic.Restore(filter)
	if err != nil {
		return helper.JSON(c, err, status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Login(c *fiber.Ctx) error {
	type jsonData struct {
		Data PortalUserLoginPayload `json:"data" validate:"required"`
	}
	var payload jsonData
	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return helper.JSON(c, parseError.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	results, status, err := Logic.Login(payload.Data)
	if err != nil {
		return helper.JSON(c, err, status)
	}

	return helper.JSON(c, results, status)
}

// Forgot

// Reset
