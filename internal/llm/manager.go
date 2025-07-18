package llm

import (
	"GoAI/internal/config"
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	
)

// StreamingCallback 定义了用于处理流式响应的回调函数类型
type StreamingCallback func(ctx context.Context, chunk []byte) error

// Manager 负责管理 LLM 客户端
type Manager struct {
	llm llms.Model
}

// NewManager 创建一个新的 LLM 管理器
func NewManager(cfg *config.AIConfig) (*Manager, error) {
	llm, err := newOpenAI(cfg)
	if err != nil {
		return nil, err
	}
	return &Manager{llm: llm}, nil
}

// GenerateContent 使用 LLM 生成内容，支持流式和非流式
func (m *Manager) GenerateContent(ctx context.Context, messages []llms.MessageContent, temperature float64, maxTokens int, callback StreamingCallback) (string, error) {
	options := []llms.CallOption{
		llms.WithTemperature(temperature),
		llms.WithMaxTokens(maxTokens),
	}

	if callback != nil {
		options = append(options, llms.WithStreamingFunc(callback))
	}

	response, err := m.llm.GenerateContent(ctx, messages, options...)
	if err != nil {
		return "", fmt.Errorf("生成内容失败: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("模型未返回任何内容")
	}

	return response.Choices[0].Content, nil
}

