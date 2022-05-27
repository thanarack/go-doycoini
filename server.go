package main

import (
	"github.com/gofiber/fiber/v2"
	"koicoini.com/sdk/v2/src/routes"
)

var router routes.Export

func main() {
	app := fiber.New()

	// Group route in v1
	api := app.Group("/api/v1")
	router.Routes(api)

	// Start server
	app.Listen(":4000")
}
