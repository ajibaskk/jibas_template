package main

import (
	"jibas-template/config"
	"jibas-template/internal/di"
	"jibas-template/internal/domain"
	"jibas-template/middleware"
	"jibas-template/pkg/swagger"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title MyApp API
// @version 1.0
// @description This is the API documentation for MyApp.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
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

	// Set up Swagger documentation route
	swagger.SetupSwagger(router)

	// Create an internal route group with JWT middleware
	internal := router.Group("/api")
	internal.Use(middleware.JWTAuthMiddleware()) // Apply JWT middleware to /internal routes

	// Initialize UserHandler and register routes within the /internal group
	userHandler := di.InitializeUserHandler(db)
	userHandler.RegisterRoutes(internal) // Pass *gin.RouterGroup

	// Run the server
	log.Fatal(router.Run(":8080"))
}
