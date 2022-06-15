package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mcuadros/go-defaults"
)

func BodyParser[T any](c *fiber.Ctx) (*T, error) {
	var req *T
	if err := c.BodyParser(&req); err != nil {
		return nil, c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}
	defaults.SetDefaults(req)
	if validator := NewValidator().Validate(req); len(validator) != 0 {
		return nil, c.Status(fiber.ErrBadRequest.Code).JSON(validator)
	}

	return req, nil
}
