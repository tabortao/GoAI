package core

import (
	"GoAI/internal/llm"
	"context"
	"github.com/tmc/langchaingo/llms"
	"io"
	"log/slog"
)

// Service encapsulates the core business logic.
type Service struct {
	llmManager *llm.LLMManager
	logger     *slog.Logger
}

// NewService creates a new core Service.
func NewService(llmManager *llm.LLMManager, logger *slog.Logger) *Service {
	return &Service{
		llmManager: llmManager,
		logger:     logger,
	}
}

// Generate performs text generation based on a prompt.
func (s *Service) Generate(ctx context.Context, prompt, modelName string, stream bool, writer io.Writer) (string, error) {
	var llmModel llms.Model
	var err error

	if modelName != "" {
		llmModel, err = s.llmManager.GetLLM(modelName)
		if err != nil {
			s.logger.Error("failed to get specified llm", "model", modelName, "error", err)
			return "", err
		}
	} else {
		llmModel, err = s.llmManager.GetDefaultLLM()
		if err != nil {
			s.logger.Error("failed to get default llm", "error", err)
			return "", err
		}
	}

	s.logger.Info("generating text", "prompt", prompt, "model", modelName, "stream", stream)

	var streamOption llms.CallOption
	if stream {
		streamOption = llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			if _, err := writer.Write(chunk); err != nil {
				return err
			}
			return nil
		})
		// When streaming, the final result is returned through the writer, so the return value of Generate is empty.
		_, err := llms.GenerateFromSinglePrompt(ctx, llmModel, prompt, streamOption)
		return "", err
	}

	// Not streaming
	completion, err := llms.GenerateFromSinglePrompt(ctx, llmModel, prompt)
	if err != nil {
		s.logger.Error("failed to generate text", "error", err)
		return "", err
	}

	return completion, nil
}
