package swagger

import (
	_ "jibas-template/docs" // Import generated Swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupSwagger configures the Swagger route for the Gin router
func SetupSwagger(router *gin.Engine) {
	// Set up the Swagger endpoint at /swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
