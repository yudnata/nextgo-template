package auth

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

func (h *Handler) Register(c fiber.Ctx) error {
	var req RegisterReq
	if err := c.Bind().Body(&req); err != nil {
		return response.JSON(c, 400, false, "Invalid body", nil)
	}
	user, err := h.service.Register(req)
	if err != nil {
		return response.JSON(c, 409, false, err.Error(), nil)
	}
	return response.JSON(c, 201, true, "Success", user)
}

func (h *Handler) Login(c fiber.Ctx) error {
	var req LoginReq
	if err := c.Bind().Body(&req); err != nil {
		return response.JSON(c, 400, false, "Invalid body", nil)
	}
	token, err := h.service.Login(req)
	if err != nil {
		return response.JSON(c, 401, false, err.Error(), nil)
	}
	return response.JSON(c, 200, true, "Success", fiber.Map{"token": token})
}

func (h *Handler) Me(c fiber.Ctx) error {
	user := c.Locals("user")
	return response.JSON(c, 200, true, "Success", user)
}
