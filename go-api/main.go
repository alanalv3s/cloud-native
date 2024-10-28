package main

import (
	"os"

	"github.com/alanalv3s/cloud-native/go-api/database"
	"github.com/alanalv3s/cloud-native/go-api/middlewares"
	"github.com/alanalv3s/cloud-native/go-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
)

func main() {
	// Initialize logger and database
	envy.Load()
	middlewares.InitLogger()
	database.InitDB()

	// Create a new Gin router
	router := gin.New()

	// Middleware: Logger and Recovery
	router.Use(gin.Recovery())                 // Gin recovery middleware
	router.Use(middlewares.LoggerMiddleware()) // Zap logger middleware

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	port := getEnv("PORT", "8080")
	if err := router.Run(":" + port); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}

// getEnv returns environment variable or fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
