package routes

import (
	"github.com/gofiber/fiber/v2"
	"koicoini.com/sdk/v2/src/controllers"
)

type Export struct{}

var controller controllers.Export

func (c Export) Routes(api fiber.Router) {
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	api.Get("/tokens", controller.TokensControllers)
}
