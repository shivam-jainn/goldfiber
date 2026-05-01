package main

import (
	"github.com/shivam-jainn/goldfiber/internal/app"
	"github.com/shivam-jainn/goldfiber/internal/config"
)

func main() {
	app := app.New()
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	app.Fiber.Listen(":" + cfg.Port)
}
