package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns a simple health check status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Service is running",
	})
}
