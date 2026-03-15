package response

import "github.com/gofiber/fiber/v3"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSON(c fiber.Ctx, status int, success bool, message string, data any) error {
	return c.Status(status).JSON(Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}
