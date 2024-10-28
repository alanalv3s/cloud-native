package routes

import (
	"github.com/alanalv3s/cloud-native/go-api/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the API routes for the application
func RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.PATCH("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	router.GET("/health", controllers.HealthCheck)
}
