package llm

import (
	"errors"
	"fmt"

	"GoAI/internal/config"
	"github.com/tmc/langchaingo/llms"
)

// LLMManager manages the lifecycle of different LLM clients.
type LLMManager struct {
	llms            map[string]llms.Model
	defaultProvider string
}

// NewLLMManager creates a new LLMManager and initializes clients based on the config.
func NewLLMManager(cfg config.Config) (*LLMManager, error) {
	manager := &LLMManager{
		llms: make(map[string]llms.Model),
	}

	// --- Unified OpenAI/Compatible API Client Initialization ---
	// Prioritize generic AI_* variables, falling back to OPENAI_*.
	providerName := cfg.AIProvider
	token := cfg.AIToken
	baseURL := cfg.AIAPIURL
	model := cfg.AIModel

	if providerName == "" {
		providerName = "openai"
	}
	if token == "" {
		token = cfg.OpenAIToken
		if token == "" {
			token = cfg.OpenAIAPIKey
		}
	}
	if baseURL == "" {
		baseURL = cfg.OpenAIBaseURL
	}
	if model == "" {
		model = cfg.OpenAIModel
	}

	// Initialize the OpenAI-compatible client if a token is available
	if token != "" {
		llm, err := NewOpenAIClient(token, baseURL, model)
		if err != nil {
			return nil, fmt.Errorf("failed to create client for provider '%s': %w", providerName, err)
		}
		manager.llms[providerName] = llm
	}

	// --- Ollama Client Initialization ---
	if cfg.OllamaURL != "" {
		ollamaLLM, err := NewOllamaClient(cfg.OllamaURL)
		if err != nil {
			return nil, fmt.Errorf("failed to create ollama client: %w", err)
		}
		manager.llms["ollama"] = ollamaLLM
	}

	if len(manager.llms) == 0 {
		return nil, errors.New("no llm providers configured")
	}

	// --- Determine Default Provider ---
	// Priority: AI_PROVIDER > "openai" > "ollama" > first available.
	if p := cfg.AIProvider; p != "" && manager.llms[p] != nil {
		manager.defaultProvider = p
	} else if manager.llms["openai"] != nil {
		manager.defaultProvider = "openai"
	} else if manager.llms["ollama"] != nil {
		manager.defaultProvider = "ollama"
	} else {
		for name := range manager.llms {
			manager.defaultProvider = name
			break
		}
	}

	return manager, nil
}

// GetLLM returns a specific LLM client by name.
func (m *LLMManager) GetLLM(name string) (llms.Model, error) {
	llm, ok := m.llms[name]
	if !ok {
		return nil, fmt.Errorf("llm provider '%s' not found or configured", name)
	}
	return llm, nil
}

// GetDefaultLLM returns the default LLM client based on the configuration priority.
func (m *LLMManager) GetDefaultLLM() (llms.Model, error) {
	if m.defaultProvider == "" {
		return nil, errors.New("no default llm provider available")
	}
	llm, ok := m.llms[m.defaultProvider]
	if !ok {
		return nil, fmt.Errorf("default llm provider '%s' not found or configured", m.defaultProvider)
	}
	return llm, nil
}