package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v3"
)

var logger *slog.Logger

func main() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Hello from Geohunter!")

	app := fiber.New()

	app.Get("ping", func(fc fiber.Ctx) error {
		return fc.Status(fiber.StatusOK).JSON("pong")
	})

	//api := app.Group("api")
	//api.Post("", nil)

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(fmt.Sprintf("Startup error: %s", err.Error()))
	}
}
