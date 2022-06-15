package client

import (
	"go-stater/domain/model"
	"go-stater/pkg/utils"

	"github.com/gofiber/fiber/v2"
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
	req, err := utils.BodyParser[model.Client](c)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}
	return c.JSON(req)
}
