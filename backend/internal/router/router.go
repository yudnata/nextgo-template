package router

import (
	"backend/internal/modules/auth"
	"backend/internal/modules/feature2"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App, authH *auth.Handler, f2H *feature2.Handler) {
	api := app.Group("/api")

	// Register Features
	auth.RegisterRoutes(api, authH)
	feature2.RegisterRoutes(api, f2H)
}
