package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"koicoini.com/sdk/v2/src/services"
)

var connection services.Export

func (c Export) TokensControllers(ctf *fiber.Ctx) error {

	tokens, err := connection.ListTokens()

	if err != nil {
		log.Panic(err)
	}

	data, _ := c.Response(tokens, "200", "Get tokens successfully")

	ctf.SendStatus(200)
	return ctf.Send(data)
}
