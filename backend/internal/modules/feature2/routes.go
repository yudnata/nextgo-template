package feature2

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router, h *Handler) {
	f2 := router.Group("/feature2")
	f2.Post("/", h.Create)
}
