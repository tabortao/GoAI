package cli

import (
	"GoAI/internal/config"
	"GoAI/internal/llm"
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
)

// streamCallback 是一个简单的回调函数，用于将流式响应实时打印到控制台
func streamCallback(ctx context.Context, chunk []byte) error {
	fmt.Print(string(chunk))
	return nil
}

// Run 启动交互式聊天 CLI
func Run(cfg *config.AIConfig) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("欢迎来到 GoAI 聊天模式! (使用模型: %s) 输入 'exit' 退出。\n", cfg.Model)

	llmManager, err := llm.NewManager(cfg)
	if err != nil {
		return fmt.Errorf("无法创建LLM管理器: %w", err)
	}

	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}
		if input == "" {
			continue
		}

		messages := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeHuman, input),
		}

		fmt.Println("AI:")
		_, err := llmManager.GenerateContent(context.Background(), messages, cfg.Temperature, 4096, streamCallback)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n生成内容时出错: %v\n", err)
		}
		fmt.Println() // 在AI响应后换行
	}

	return nil
}

// Generate 执行一次性内容生成
func Generate(cfg *config.AIConfig, prompt string) error {
	llmManager, err := llm.NewManager(cfg)
	if err != nil {
		return fmt.Errorf("无法创建LLM管理器: %w", err)
	}

	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "你是一个通用AI助手，负责根据用户的请求生成内容。"),
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	}

	fmt.Println("正在生成内容...")

	_, err = llmManager.GenerateContent(context.Background(), messages, cfg.Temperature, 4096, streamCallback)
	if err != nil {
		return fmt.Errorf("\n生成内容时出错: %w", err)
	}
	fmt.Println() // 在所有内容生成后换行
	return nil
}