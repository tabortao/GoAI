package llm

import (
	"GoAI/internal/config"

	"github.com/tmc/langchaingo/llms/openai"
)

func newOpenAI(cfg *config.AIConfig) (*openai.LLM, error) {
	llm, err := openai.New(
		openai.WithModel(cfg.Model),
		openai.WithToken(cfg.Token),
		openai.WithBaseURL(cfg.URL),
	)
	if err != nil {
		return nil, err
	}
	return llm, nil
}