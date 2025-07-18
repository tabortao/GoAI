package cli

import (
	"GoAI/internal/config"
	"GoAI/internal/llm"
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
)

var modelName string

func init() {
	chatCmd.Flags().StringVarP(&modelName, "model", "m", "", "指定要使用的AI模型名称")
	rootCmd.AddCommand(chatCmd)
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "启动交互式聊天模式",
	Long:  `与指定的AI模型进行交互式聊天。输入 'exit' 退出。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig("config.json")
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		modelConfig, err := cfg.GetModelConfig(modelName)
		if err != nil {
			return err
		}

		llmManager, err := llm.NewManager(modelConfig)
		if err != nil {
			return fmt.Errorf("无法创建LLM管理器: %w", err)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("欢迎来到 GoAI 聊天模式! (使用模型: %s) 输入 'exit' 退出。\n", modelConfig.Model)

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
			_, err := llmManager.GenerateContent(context.Background(), messages, modelConfig.Temperature, 4096, streamCallback)
			if err != nil {
				fmt.Fprintf(os.Stderr, "\n生成内容时出错: %v\n", err)
			}
			fmt.Println() // 在AI响应后换行
		}
		return nil
	},
}

// streamCallback 是一个简单的回调函数，用于将流式响应实时打印到控制台
func streamCallback(ctx context.Context, chunk []byte) error {
	fmt.Print(string(chunk))
	return nil
}
