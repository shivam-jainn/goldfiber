package handler

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// ErrorResponse represents the structure of the errored API response
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

// SuccessResponse represents the structure of a successful API response
type SuccessResponse struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// ErrorHandler is a centralized error handler for Fiber
func ErrorHandler(c fiber.Ctx, err error) error {
	// Default 500 status code
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Retreive the custom status code if it's a fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// Log the error
	log.Printf("Error: %v", err)
	// If you use zap logger from internal/logger: logger.Log.Error(err.Error())

	// Send JSON response
	return c.Status(code).JSON(ErrorResponse{
		Error:   true,
		Message: message,
		Code:    code,
	})
}

// SendSuccess is a helper function to send successful JSON responses
func SendSuccess(c fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(SuccessResponse{
		Error: false,
		Data:  data,
	})
}

// SendError is a helper function to send error JSON responses manually
func SendError(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(ErrorResponse{
		Error:   true,
		Message: message,
		Code:    status,
	})
}
