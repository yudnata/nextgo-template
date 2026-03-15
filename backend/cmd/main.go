package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/modules/auth"
	"backend/internal/modules/feature2"
	"backend/internal/router"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)
	database.Migrate(db)

	// Auth Module
	authRepo := auth.NewRepository(db)
	authServ := auth.NewService(authRepo, cfg)
	authHand := auth.NewHandler(authServ)

	// Feature2 Module
	f2Repo := feature2.NewRepository(db)
	f2Serv := feature2.NewService(f2Repo)
	f2Hand := feature2.NewHandler(f2Serv)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	router.Setup(app, authHand, f2Hand)

	log.Fatal(app.Listen(":" + cfg.Port))
}
