package app

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	})

	// Core Middlewares
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.FiberMiddleware())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
	}))

	// Rate limiter middleware
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
	}))

	// Metrics endpoint (Prometheus)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
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
