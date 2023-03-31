package company

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	md "app/cmd/models"
	"app/pkg/helper"
)

func (h *handler) List(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.CompanyFilter `json:"filter"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return helper.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	results, status, err := Logic.List(payload.Filter)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, results, status)
}

func (h *handler) Insert(c *fiber.Ctx) error {
	type jsonData struct {
		Data md.Company `json:"data"`
	}
	var payload jsonData

	err := c.BodyParser(&payload)
	if err != nil {
		return helper.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	company := payload.Data
	result, status, err := Logic.Insert(company)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Update(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.CompanyFilter `json:"filter"`
		Data   md.CompanyUpdate `json:"data"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return helper.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	company := payload.Data
	result, status, err := Logic.Update(filter, company)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Archive(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.CompanyFilter `json:"filter"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return helper.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	result, status, err := Logic.Archive(filter)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Restore(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.CompanyFilter `json:"filter"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return helper.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := helper.Validator(payload)
	if validationError.Errors != nil {
		return helper.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	result, status, err := Logic.Restore(filter)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, result, status)
}
