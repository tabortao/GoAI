package llm

import (
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// NewOpenAIClient creates a new OpenAI LLM client.
func NewOpenAIClient(token, baseURL, model string) (llms.Model, error) {
	opts := []openai.Option{
		openai.WithToken(token),
	}

	if baseURL != "" {
		opts = append(opts, openai.WithBaseURL(baseURL))
	}

	if model != "" {
		opts = append(opts, openai.WithModel(model))
	}

	return openai.New(opts...)
}
