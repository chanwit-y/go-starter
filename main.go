package main

import (
	routers "go-stater/router"

	_ "go-stater/docs"

	"github.com/gofiber/fiber/v2"
)

// @title Banpu API
// @version 1.0.0
// @description API Endpoint for BanpuApps
// @termsOfService
// @contact.name Applicatoin Development
// @contact.email fiber@swagger.io
// @license.name Banpu Co.ltd.
// @host localhost:3001
// @BasePath /api/v1/
func main() {
	app := fiber.New()
	routers.PublicRoutes(app)

	app.Listen(":3000")

}
