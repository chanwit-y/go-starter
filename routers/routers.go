package routers

import (
	"go-stater/module/master/client"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	routerV1 := app.Group("v1")
	client.Router(routerV1)
}
