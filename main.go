package main

import (
	"went-template/internal/config"
	"went-template/internal/middlewares"
	"went-template/internal/providers"
	"went-template/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configurations
	config := &config.Config{}
	config.Initialize()

	// Initialize database connection
	dbp := &providers.DatabaseProvider{}
	dbp.Connect(config)

	// Initialize Gin router
	r := gin.Default()

	// Middlewares
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middlewares.Cors(),
		middlewares.Auth(config.JWTSecret, dbp.DB),
	)

	// Setup routes
	routes.SetupRoutes(r, dbp.DB)

	// Start the server
	r.Run(":" + config.Port)
}
