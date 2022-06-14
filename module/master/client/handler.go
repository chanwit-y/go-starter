package client

import (
	"go-stater/domain/model"
	"go-stater/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/mcuadros/go-defaults"
)

type handler struct {
	service Service
}

func NewHandler(service Service) handler {
	return handler{service: service}
}

func (h *handler) FindAll(c *fiber.Ctx) error {
	return c.JSON(h.service.FindAll())
}

func (h *handler) Create(c *fiber.Ctx) error {
	var req *model.Client
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}
	defaults.SetDefaults(req)
	if validator := utils.NewValidator().Validate(req); len(validator) != 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(validator)
	}

	return c.JSON(req)
}
