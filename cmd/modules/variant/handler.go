package variant

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	md "app/cmd/models"
	"app/pkg/helper"
)

func (h *handler) List(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.VariantFilter `json:"filter"`
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
		Data md.Variant `json:"data"`
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

	variant := payload.Data
	result, status, err := Logic.Insert(variant)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Update(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.VariantFilter `json:"filter"`
		Data   md.VariantUpdate `json:"data"`
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
	variant := payload.Data
	result, status, err := Logic.Update(filter, variant)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Archive(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.VariantFilter `json:"filter"`
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
		Filter md.VariantFilter `json:"filter"`
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
