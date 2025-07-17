package cli

import (
	"GoAI/internal/api"
	"GoAI/internal/config"
	"GoAI/internal/core"
	"GoAI/internal/llm"
	"GoAI/internal/models"
	"GoAI/pkg/utils"
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	stream   bool
	modelName string
	text      string
	template  string
)

// NewRootCmd creates the root command for the CLI application.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "goai",
		Short: "GoAI is a versatile AI text processing tool.",
		Long:  `A dual-mode application that provides both a RESTful API and a command-line interface for interacting with Large Language Models.`,
	}

	rootCmd.AddCommand(newServerCmd())
	rootCmd.AddCommand(newGenerateCmd())
	rootCmd.AddCommand(newChatCmd())

	return rootCmd
}

func newServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start the HTTP API server",
		Run:   runServer,
	}
}

func newGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate [prompt]",
		Short: "Generate text from a single prompt",
		Args:  cobra.ExactArgs(1),
		Run:   runGenerate,
	}
	cmd.Flags().BoolVar(&stream, "stream", false, "Enable streaming output")
	cmd.Flags().StringVar(&modelName, "model", "", "Specify the model to use (e.g., openai, ollama)")
	cmd.Flags().StringVar(&text, "text", "", "Add text to the prompt")
	cmd.Flags().StringVar(&template, "template", "", "Specify a template to use")
	return cmd
}

func newChatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat",
		Short: "Start an interactive chat session",
		Run:   runChat,
	}
	cmd.Flags().StringVar(&modelName, "model", "", "Specify the model to use for the chat session (e.g., openai, ollama)")
	return cmd
}

func runServer(cmd *cobra.Command, args []string) {
	cfg, logger := setup()

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

func runGenerate(cmd *cobra.Command, args []string) {
	cfg, logger := setup()

	llmManager, err := llm.NewLLMManager(cfg)
	if err != nil {
		logger.Error("failed to create llm manager", "error", err)
		os.Exit(1)
	}

	service := core.NewService(llmManager, logger)
	req := &models.GenerateRequest{
		Prompt:   args[0],
		Text:     text,
		Template: template,
		Stream:   stream,
		Model:    modelName,
	}

	_, err = service.Generate(context.Background(), req, os.Stdout)
	if err != nil {
		logger.Error("failed to generate text", "error", err)
		os.Exit(1)
	}
	fmt.Println() // Add a newline for better formatting
}

func runChat(cmd *cobra.Command, args []string) {
	cfg, logger := setup()

	llmManager, err := llm.NewLLMManager(cfg)
	if err != nil {
		logger.Error("failed to create llm manager", "error", err)
		os.Exit(1)
	}

	service := core.NewService(llmManager, logger)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Starting interactive chat session. Type 'exit' or 'quit' to end.")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		prompt := scanner.Text()
		if prompt == "exit" || prompt == "quit" {
			break
		}

		req := &models.GenerateRequest{
			Prompt: prompt,
			Stream: true,
			Model:  modelName,
		}

		_, err := service.Generate(context.Background(), req, os.Stdout)
		if err != nil {
			logger.Error("failed to generate response", "error", err)
		}
		fmt.Println()
	}
}

func setup() (config.Config, *slog.Logger) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	logger := utils.NewLogger(cfg.LogLevel)
	return cfg, logger
}
