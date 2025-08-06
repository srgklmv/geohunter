package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/srgklmv/geohunter/internal/app"
	"github.com/srgklmv/geohunter/internal/controller"
)

var logger *slog.Logger

func main() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Hello from Geohunter!")

	app := app.New()
	controller.SetupRoutes(app)

	if err := app.Listen("localhost:3000"); err != nil {
		panic(fmt.Sprintf("Startup error: %s", err.Error()))
	}
}
