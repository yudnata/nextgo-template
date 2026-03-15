package feature2

import (
	"backend/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(c fiber.Ctx) error {
	var req CreateReq
	if err := c.Bind().Body(&req); err != nil {
		return response.JSON(c, 400, false, "Invalid body", nil)
	}
	res, err := h.service.Create(req)
	if err != nil {
		return response.JSON(c, 500, false, err.Error(), nil)
	}
	return response.JSON(c, 201, true, "Success", res)
}
