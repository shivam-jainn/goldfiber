package handler

import "github.com/gofiber/fiber/v3"

func getHealth(c fiber.Ctx) error {
	return c.SendString("OK")
}

func getLiveness(c fiber.Ctx) error {
	return SendSuccess(c, fiber.StatusOK, map[string]string{
		"status": "alive",
	})
}

func getReadiness(c fiber.Ctx) error {
	return SendSuccess(c, fiber.StatusOK, map[string]string{
		"status": "ready",
	})
}

func getDetailedHealth(c fiber.Ctx) error {
	detailedHealth := map[string]interface{}{
		"status": "OK",
	}
	return SendSuccess(c, fiber.StatusOK, detailedHealth)
}

func RegisterHealthRoutes(router fiber.Router) {
	healthRouter := router.Group("/health")

	healthRouter.Get("/", getHealth)
	healthRouter.Get("/liveness", getLiveness)
	healthRouter.Get("/readiness", getReadiness)
	healthRouter.Get("/details", getDetailedHealth)
}
