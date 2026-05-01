package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/shivam-jainn/goldfiber/internal/handler"
	"github.com/shivam-jainn/goldfiber/internal/logger"
)

type App struct {
	Fiber *fiber.App
}

func New() *App {
	app := fiber.New()

	app.Use(logger.FiberMiddleware())

	// routes
	handler.RegisterHealthRoutes(app)

	return &App{Fiber: app}
}
