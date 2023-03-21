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
	portal_user := new(PortalUser)
	parseError := c.BodyParser(portal_user)
	if parseError != nil {
		return helper.JSON(c, parseError.Error(), http.StatusBadRequest)
	}

	result, status, logicError := Logic.Insert(*portal_user)
	if logicError != nil {
		return helper.JSON(c, logicError, status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Update(c *fiber.Ctx) error {
	type JsonData struct {
		Filter     PortalUserFilter `json:"filter"`
		PortalUser PortalUser       `json:"portal_user"`
	}
	var payload JsonData
	parseError := c.BodyParser(&payload)
	if parseError != nil {
		return helper.JSON(c, parseError.Error(), http.StatusBadRequest)
	}

	filter := payload.Filter
	portal_user := payload.PortalUser
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
