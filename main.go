package main

import (
	"fmt"
	"log"
	"uni_app/database"
	"uni_app/pkg"

	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	config := env.NewConfig()

	// Connect to database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize database
	if err := database.Init(db); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// // Initialize repositories
	// authRepo := repositories.NewAuthRepository(db)

	// // Initialize usecases
	// authUsecase := usecase.NewAuthUsecase(authRepo, config)

	// // Initialize middleware
	// _ = middleware.NewAuthMiddleware(authRepo, config)

	// Initialize Echo
	e := echo.New()
	apiGroup := e.Group("/api")
	pkg.InitPkgs(db, apiGroup, config)

	// Initialize handlers
	// handler.NewAuthHandler(authUsecase, *apiGroup)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
