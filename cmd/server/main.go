package main

import (
	"GoAI/internal/api"
	"GoAI/internal/config"
	"GoAI/internal/core"
	"GoAI/internal/llm"
	"GoAI/pkg/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := utils.NewLogger(cfg.LogLevel)

	llmManager, err := llm.NewLLMManager(cfg)
	if err != nil {
		logger.Error("failed to create llm manager", "error", err)
		os.Exit(1)
	}

	service := core.NewService(llmManager, logger)
	handler := api.NewAPIHandler(service, logger)
	router := api.SetupRouter(handler)

	srv := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: router,
	}

	go func() {
		logger.Info("starting server", "port", cfg.HTTPPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}
	logger.Info("server exited gracefully")
}
