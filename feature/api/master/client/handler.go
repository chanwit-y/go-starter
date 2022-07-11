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

// Application godoc
// @Summary      Get all
// @Description  get client info 1
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param	 limit      query    number     false    "limit"
// @Param	 offset     query    number     false    "offset"
// @Success      200 {string} string "ok"
// @Router       /client [get]
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
