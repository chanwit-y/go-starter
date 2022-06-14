package main

import (
	"go-stater/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routers.PublicRoutes(app)

	app.Listen(":3000")

}
