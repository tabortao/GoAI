package cli

import (
	"GoAI/internal/config"
	"GoAI/internal/server"
	"fmt"

	"github.com/spf13/cobra"
)

var port string

func init() {
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "指定API服务器的端口")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动API服务器",
	Long:  `启动一个HTTP服务器，提供用于生成内容的API端点。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig("config.json")
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		fmt.Printf("启动API服务器在端口 %s\n", port)
		apiServer := server.NewServer(cfg)
		if err := apiServer.Run(":" + port); err != nil {
			return fmt.Errorf("无法启动服务器: %w", err)
		}
		return nil
	},
}

