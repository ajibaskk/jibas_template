package main

import (
	"jibas-template/config"
	"jibas-template/internal/di"
	"jibas-template/internal/domain"
	"jibas-template/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file if it exists
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, falling back to system environment variables")
	}
	// Initialize the database
	db := config.InitDB()
	defer func() {
		dbConn, err := db.DB()
		if err == nil {
			dbConn.Close()
		}
	}()

	// Automatically migrate the User model
	db.AutoMigrate(&domain.User{})

	// Initialize Gin router
	router := gin.Default()

	// Create an internal route group with JWT middleware
	internal := router.Group("/api")
	internal.Use(middleware.JWTAuthMiddleware()) // Apply JWT middleware to /internal routes

	// Initialize UserHandler and register routes within the /internal group
	userHandler := di.InitializeUserHandler(db)
	userHandler.RegisterRoutes(internal) // Pass *gin.RouterGroup

	// Run the server
	log.Fatal(router.Run(":8080"))
}
