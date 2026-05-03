package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/shivam-jainn/goldfiber/internal/app"
	"github.com/shivam-jainn/goldfiber/internal/config"
	"github.com/shivam-jainn/goldfiber/internal/database"
	"github.com/shivam-jainn/goldfiber/internal/logger"
)

func main() {
	// Load config
	cfg := &config.Config{}
	if err := config.LoadConfig(cfg); err != nil {
		panic(err)
	}

	// Setup logger
	logger.SetLogger(logger.NewZap(cfg))

	logger.Info("Configuration loaded",
		logger.Field{Key: "env", Value: cfg.Env},
		logger.Field{Key: "port", Value: cfg.Port},
	)

	// Connect DB
	pool, err := database.Connect(*cfg)
	if err != nil {
		logger.Error("DB connection failed", logger.Field{Key: "error", Value: err.Error()})
		panic(err)
	}
	defer pool.Close() // ✅ correct place

	// Create app with dependencies
	a := app.New(cfg, pool)

	// Run server in goroutine
	go func() {
		logger.Info("Starting server...",
			logger.Field{Key: "port", Value: cfg.Port},
		)

		if err := a.Fiber.Listen(":" + cfg.Port); err != nil {
			logger.Error("Server error", logger.Field{Key: "error", Value: err.Error()})
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Shutting down server...")

	if err := a.Fiber.Shutdown(); err != nil {
		logger.Error("Shutdown error", logger.Field{Key: "error", Value: err.Error()})
	}

	logger.Info("Server stopped cleanly")
}
