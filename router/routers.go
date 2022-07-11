package routers

import (
	"go-stater/feature/api/master/client"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func PublicRoutes(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault) // default

	routerV1 := app.Group("v1")
	client.Router(routerV1)

}
