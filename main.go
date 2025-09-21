package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
		// Add other middlewares as needed
	)

	// Setup routes
	routes.SetupRoutes(r, dbp.DB)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: r,
	}

	// Create context for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", config.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()
	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
