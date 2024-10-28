package middlewares

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// InitLogger initializes the Zap logger.
func InitLogger() {

	var zapLogger *zap.Logger
	var err error

	if os.Getenv("ENV") == "PRODUCTION" {
		zapLogger, err = zap.NewProduction()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	logger = zapLogger.Sugar()
	defer zapLogger.Sync() // Flush any buffered log entries
}

// LoggerMiddleware logs incoming requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // Process the request

		latency := time.Since(start)
		logger.Infow("Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency,
			"clientIP", c.ClientIP(),
		)
	}
}
