package client

import "github.com/gofiber/fiber/v2"

func Router(r fiber.Router) {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	groupRoute := r.Group("/client")
	groupRoute.Get("", handler.FindAll)
	groupRoute.Post("", handler.Create)
}
