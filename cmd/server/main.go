package main

import (
	"github.com/shivam-jainn/goldfiber/internal/app"
)

func main() {
	app := app.New()

	app.Fiber.Listen(":8080")
}
