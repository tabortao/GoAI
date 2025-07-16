package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the API routes.
func SetupRouter(handler *APIHandler) *gin.Engine {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", handler.HealthCheckHandler)

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		v1.POST("/generate", handler.GenerateHandler)
	}

	return r
}
