package handler

import "github.com/gofiber/fiber/v3"

func getHealth(c fiber.Ctx) error {
	return c.SendString("OK")
}

func getDetailedHealth(c fiber.Ctx) error {
	detailedHealth := map[string]interface{}{
		"status": "OK",
	}
	return c.JSON(detailedHealth)
}

func RegisterHealthRoutes(router fiber.Router) {
	healthRouter := router.Group("/health")

	healthRouter.Get("/", getHealth)
	healthRouter.Get("/details", getDetailedHealth)
}
