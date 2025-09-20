package routes

import (
	"went-template/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	// Basic Health Check Endpoint
	r.GET("/ping", HealthCheck)

	// Swagger setup
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User routes
	UserRoutes(r, db)
	// [*RP*] Please do not delete this comment. It is used to automatically add new route files.
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "pong",
	})
}
