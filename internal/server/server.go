package server

import (
	"GoAI/internal/api"
	"GoAI/internal/config"

	"github.com/gin-gonic/gin"
)

// Server struct
type Server struct {
	router *gin.Engine
	config *config.Config
}

// NewServer creates a new server
func NewServer(cfg *config.Config) *Server {
	if cfg.GinMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.Default()
	server := &Server{
		router: router,
		config: cfg,
	}

	api.RegisterRoutes(router, cfg) // 注入配置

	return server
}

// Run starts the server
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
