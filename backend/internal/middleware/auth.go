package middleware

import (
	"strings"

	"backend/internal/auth/repository"
	"backend/internal/auth/service"
	"backend/internal/common/model"

	"github.com/gofiber/fiber/v3"
)

func AuthRequired(authService *service.AuthService, userRepo *repository.UserRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(model.Response{Success: false, Message: "Missing authorization header"})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return c.Status(401).JSON(model.Response{Success: false, Message: "Invalid authorization format"})
		}

		userID, err := authService.ValidateToken(parts[1])
		if err != nil {
			return c.Status(401).JSON(model.Response{Success: false, Message: "Invalid or expired token"})
		}

		user, err := userRepo.FindByID(c.Context(), userID)
		if err != nil {
			return c.Status(401).JSON(model.Response{Success: false, Message: "User not found"})
		}

		c.Locals("user", user)
		c.Locals("userID", userID)
		return c.Next()
	}
}
