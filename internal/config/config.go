package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/srgklmv/geohunter/pkg/logger"
)

type Config struct {
	Database Database `json:"database"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}

func New() (Config, error) {
	// TODO: add executable path
	dir, err := os.Getwd()
	if err != nil {
		logger.Error("new config error", slog.String("err", err.Error()))
		return Config{}, err
	}

	fmt.Print("\n REMOVE ME! ", "dir: ", dir, "\n")

	return Config{}, nil
}
