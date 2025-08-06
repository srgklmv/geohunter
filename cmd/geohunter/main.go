package main

import (
	"fmt"

	"github.com/srgklmv/geohunter/internal/app"
	"github.com/srgklmv/geohunter/internal/controller"
	"github.com/srgklmv/geohunter/pkg/logger"
)

func main() {
	logger.Init()

	logger.Info("Hello from Geohunter!")

	app := app.New()
	controller.SetupRoutes(app)

	if err := app.Listen("localhost:3000"); err != nil {
		panic(fmt.Sprintf("Startup error: %s", err.Error()))
	}
}
