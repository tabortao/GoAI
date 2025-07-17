package core

import (
	"GoAI/internal/llm"
	"GoAI/internal/models"
	"GoAI/pkg/utils"
	"context"
	"fmt"
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
func (s *Service) Generate(ctx context.Context, req *models.GenerateRequest, writer io.Writer) (string, error) {
	var llmModel llms.Model
	var err error

	if req.Model != "" {
		llmModel, err = s.llmManager.GetLLM(req.Model)
		if err != nil {
			s.logger.Error("failed to get specified llm", "model", req.Model, "error", err)
			return "", err
		}
	} else {
		llmModel, err = s.llmManager.GetDefaultLLM()
		if err != nil {
			s.logger.Error("failed to get default llm", "error", err)
			return "", err
		}
	}

	prompt, err := s.buildPrompt(req)
	if err != nil {
		return "", err
	}

	s.logger.Info("generating text", "prompt", prompt, "model", req.Model, "stream", req.Stream)

	var streamOption llms.CallOption
	if req.Stream {
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

func (s *Service) buildPrompt(req *models.GenerateRequest) (string, error) {
	fullPrompt := req.Prompt
	if req.Text != "" {
		fullPrompt = fmt.Sprintf("%s %s", req.Prompt, req.Text)
	}

	if req.Template != "" {
		data := map[string]string{
			"prompt": fullPrompt,
		}
		return utils.ApplyTemplate(req.Template, data)
	}

	return fullPrompt, nil
}
