package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goai",
	Short: "GoAI 是一个通用的AI命令行工具",
	Long:  `一个使用 Go 语言编写的、支持多种模型的 AI 命令行工具，可以用于聊天、生成内容和作为 API 服务器。`,
}

// Execute 执行根命令，这是所有 CLI 命令的入口点
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "执行命令时出错: %v\n", err)
		os.Exit(1)
	}
}

func init() {
}