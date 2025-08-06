package app

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/srgklmv/geohunter/internal/controller"
	"github.com/srgklmv/geohunter/pkg/logger"
)

type app struct {
	app *fiber.App
}

func New() *app {
	// TODO: Add config.
	return &app{
		app: fiber.New(),
	}
}

func (a *app) Run() error {
	controller.SetupRoutes(a.app)

	if err := a.app.Listen("localhost:3000"); err != nil {
		logger.Error("Server listen error.", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (a *app) Shutdown() {

}
