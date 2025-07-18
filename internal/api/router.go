package api

import (
	"GoAI/internal/config"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the API routes
func RegisterRoutes(router *gin.Engine, cfg *config.Config) {
	api := router.Group("/api/v1")
	{
		api.POST("/generate", GenerateHandler(cfg))
	}
}