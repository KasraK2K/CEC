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

	filter := payload.Filter
	results, status, logicError := Logic.List(filter)
	if logicError != nil {
		return helper.JSON(c, logicError, status)
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

	portalUser := payload.Data
	result, status, logicError := Logic.Insert(portalUser)
	if logicError != nil {
		return helper.JSON(c, logicError, status)
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

	filter := payload.Filter
	portal_user := payload.Data
	result, status, logicError := Logic.Update(filter, portal_user)
	if len(logicError) > 0 {
		return helper.JSON(c, logicError, status)
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

	filter := payload.Filter
	result, status, logicError := Logic.Archive(filter)
	if logicError != nil {
		return helper.JSON(c, logicError, status)
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

	filter := payload.Filter
	result, status, logicError := Logic.Restore(filter)
	if logicError != nil {
		return helper.JSON(c, logicError, status)
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

	results, status, logicError := Logic.Login(payload.Data)
	if logicError != nil {
		return helper.JSON(c, logicError, status)
	}

	return helper.JSON(c, results, status)
}

// Forgot

// Reset
