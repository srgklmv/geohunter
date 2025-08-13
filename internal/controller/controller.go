package controller

import (
	"github.com/gofiber/fiber/v3"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func SetupRoutes(app *fiber.App) {
	// Last part of initialization is ping.
	app.Get("ping", func(fc fiber.Ctx) error {
		return fc.Status(fiber.StatusOK).JSON("pong")
	})
}
