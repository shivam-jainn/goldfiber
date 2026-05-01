package logger

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func FiberMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		Info("http_request",
			Field{"method", c.Method()},
			Field{"path", c.Path()},
			Field{"status", c.Response().StatusCode()},
			Field{"latency", time.Since(start).String()},
			Field{"ip", c.IP()},
		)

		return err
	}
}
