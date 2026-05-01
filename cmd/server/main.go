package main

import (
	"github.com/shivam-jainn/goldfiber/internal/app"
	"github.com/shivam-jainn/goldfiber/internal/config"
	"github.com/shivam-jainn/goldfiber/internal/logger"
)

func main() {
	cfg := &config.Config{}
	err := config.LoadConfig(cfg)
	if err != nil {
		panic(err)
	}

	logger.SetLogger(logger.NewZap(cfg))

	logger.Info("Configuration loaded",
		logger.Field{Key: "env", Value: cfg.Env},
		logger.Field{Key: "debug", Value: cfg.Debug},
		logger.Field{Key: "log_level", Value: cfg.LogLevel},
	)

	logger.Info("Starting application...")
	app := app.New()

	logger.Info("Starting server...",
		logger.Field{Key: "port", Value: cfg.Port},
	)

	if err := app.Fiber.Listen(":" + cfg.Port); err != nil {
		logger.Error("Failed to start server", logger.Field{Key: "error", Value: err.Error()})
	}

	app.Fiber.Listen(":" + cfg.Port)

	logger.Info("Server stopped")
}
