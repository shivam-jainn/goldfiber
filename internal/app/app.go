package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shivam-jainn/goldfiber/internal/config"
	"github.com/shivam-jainn/goldfiber/internal/handler"
	"github.com/shivam-jainn/goldfiber/internal/logger"
)

type App struct {
	Fiber *fiber.App
	DB    *pgxpool.Pool
	Cfg   *config.Config
}

func New(cfg *config.Config, pool *pgxpool.Pool) *App {
	app := fiber.New()

	// middleware
	app.Use(logger.FiberMiddleware())

	// versioning
	// api := app.Group("/api")
	// v1 := api.Group("/v1")

	// routes
	handler.RegisterHealthRoutes(app)
	// future:
	// handler.RegisterAuthRoutes(v1.Group("/auth"))

	return &App{
		Fiber: app,
		DB:    pool,
		Cfg:   cfg,
	}
}
