package portal_user

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	md "app/cmd/models"
	"app/pkg/helper"
)

func (h *handler) List(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.PortalUserFilter `json:"filter"`
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
		Data md.PortalUser `json:"data"`
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

	portalUser := payload.Data
	result, status, err := Logic.Insert(portalUser)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Update(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.PortalUserFilter `json:"filter"`
		Data   md.PortalUserUpdate `json:"data"`
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
	update := payload.Data
	result, status, err := Logic.Update(filter, update)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, result, status)
}

func (h *handler) Archive(c *fiber.Ctx) error {
	type JsonData struct {
		Filter md.PortalUserFilter `json:"filter"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return helper.JSON(c, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(payload.Filter)

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
		Filter md.PortalUserFilter `json:"filter"`
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

func (h *handler) Login(c *fiber.Ctx) error {
	type jsonData struct {
		Data md.PortalUserLoginPayload `json:"data" validate:"required"`
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

	results, status, err := Logic.Login(payload.Data)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, results, status)
}

func (h *handler) ForgotPassword(c *fiber.Ctx) error {
	type forgetPass struct {
		Email string `json:"email" bson:"email" gorm:"type:string;unique;not null;" validate:"required,email,min=6,max=32"`
	}
	type JsonData struct {
		Filter forgetPass `json:"filter"`
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

	results, status, err := Logic.ForgotPassword(payload.Filter.Email)
	if err != nil {
		return helper.JSON(c, err.Error(), status)
	}

	return helper.JSON(c, results, status)
}

func (h *handler) Upload(c *fiber.Ctx) error {
	type upload struct {
		M string `json:"m"`
		S string `json:"s"`
	}
	var payload upload
	err := c.BodyParser(&payload)
	if err != nil {
		return helper.JSON(c, err.Error(), http.StatusBadRequest)
	}
	return helper.JSON(c, payload, http.StatusOK)
}
