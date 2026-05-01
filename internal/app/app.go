package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/shivam-jainn/goldfiber/internal/handler"
)

type App struct {
	Fiber *fiber.App
}

func New() *App {
	app := fiber.New()

	// // versioning
	// api := app.Group("/api")
	// v1 := api.Group("/v1")

	// routes
	handler.RegisterHealthRoutes(app)

	return &App{Fiber: app}
}
