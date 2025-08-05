package api

import (
	"GoAI/internal/config"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the API routes
func RegisterRoutes(router *gin.Engine, cfg *config.Config) {
	// 静态文件服务
	router.StaticFile("/", "./web/index.html")
	router.Static("/web", "./web")

	api := router.Group("/api/v1")
	{
		api.POST("/generate", GenerateHandler(cfg))
	}
}
