package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/srgklmv/geohunter/internal/app"
	"github.com/srgklmv/geohunter/pkg/logger"
)

func main() {
	logger.Init()
	logger.Info("Starting up Geohunter...")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	app := app.New()

	go func() {
		if err := app.Run(); err != nil {
			logger.Error("Startup error. Exiting Geohunter...", slog.String("error", err.Error()))
			shutdown <- syscall.SIGTERM
		}
	}()
	logger.Info("Geohunter is running.")

	<-shutdown
	logger.Info("Shutting down Geohunter...")

	app.Shutdown()
	logger.Info("Geohunter shut down.")
}
