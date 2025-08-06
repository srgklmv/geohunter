package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v3"
)

var logger *slog.Logger

func main() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Hello from Geohunter!")

	app := fiber.New()

	api := app.Group("api")
	api.Post("", nil)
}
