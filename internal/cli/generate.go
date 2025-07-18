package cli

import (
	"GoAI/internal/config"
	"GoAI/internal/llm"
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
)

var text string

func init() {
	generateCmd.Flags().StringVarP(&modelName, "model", "m", "", "指定要使用的AI模型名称")
	generateCmd.Flags().StringVarP(&text, "text", "t", "", "需要处理的附加文本内容")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate [prompt]",
	Short: "根据提示生成一次性内容",
	Long:  `根据用户提供的提示（和可选的--text标志），调用AI模型生成内容并以流式方式输出。`,
	Args:  cobra.MaximumNArgs(1), // prompt 是可选的，最多一个
	RunE: func(cmd *cobra.Command, args []string) error {
		var prompt string
		if len(args) > 0 {
			prompt = args[0]
		}

		if prompt == "" && text == "" {
			return fmt.Errorf("错误: 必须提供一个 prompt 参数或使用 --text 标志")
		}

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

		// 组合 prompt 和 text
		finalPrompt := prompt
		if text != "" {
			if finalPrompt != "" {
				finalPrompt = fmt.Sprintf("%s\n\n---\n\n%s", prompt, text)
			} else {
				finalPrompt = text
			}
		}

		messages := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, "你是一个通用AI助手，负责根据用户的请求生成内容。"),
			llms.TextParts(llms.ChatMessageTypeHuman, finalPrompt),
		}

		fmt.Println("正在生成内容...")

		_, err = llmManager.GenerateContent(context.Background(), messages, modelConfig.Temperature, 4096, streamCallback)
		if err != nil {
			return fmt.Errorf("\n生成内容时出错: %w", err)
		}
		fmt.Println() // 在所有内容生成后换行
		return nil
	},
}