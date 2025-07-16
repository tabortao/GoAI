package llm

import (
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// NewOllamaClient creates a new Ollama LLM client.
func NewOllamaClient(url string) (llms.Model, error) {
	return ollama.New(ollama.WithServerURL(url))
}
