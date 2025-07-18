package main

import (
	"GoAI/internal/cli"
	"GoAI/internal/config"
	"GoAI/internal/server"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// 加载配置
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: 加载配置失败: %v\n", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		handleGenerate(cfg)
	case "chat":
		handleChat(cfg)
	case "server":
		handleServer(cfg)
	default:
		printUsage()
	}
}

func handleGenerate(cfg *config.Config) {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	modelName := generateCmd.String("model", "", "要使用的AI模型名称")

	if len(os.Args) < 3 {
		fmt.Println("错误: 'generate' 命令需要一个提示参数")
		generateCmd.Usage()
		os.Exit(1)
	}

	// 手动解析 os.Args[2:]
	// 我们需要分离出标志和参数
	var prompt string
	args := os.Args[2:]
	// 解析标志
	generateCmd.Parse(args)

	// 提取提示文本（非标志参数）
	if generateCmd.NArg() > 0 {
		prompt = strings.Join(generateCmd.Args(), " ")
	} else {
		fmt.Println("错误: 'generate' 命令需要一个提示参数")
		generateCmd.Usage()
		os.Exit(1)
	}

	modelConfig, err := cfg.GetModelConfig(*modelName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	if err := cli.Generate(modelConfig, prompt); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}

func handleChat(cfg *config.Config) {
	chatCmd := flag.NewFlagSet("chat", flag.ExitOnError)
	modelName := chatCmd.String("model", "", "要使用的AI模型名称")
	chatCmd.Parse(os.Args[2:])

	modelConfig, err := cfg.GetModelConfig(*modelName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	if err := cli.Run(modelConfig); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}

func handleServer(cfg *config.Config) {
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	// 可以为server添加特定参数，例如 --port
	port := serverCmd.String("port", "8080", "API服务器端口")
	serverCmd.Parse(os.Args[2:])

	fmt.Printf("启动API服务器在端口 %s\n", *port)
	apiServer := server.NewServer(cfg)
	if err := apiServer.Run(":" + *port); err != nil {
		fmt.Fprintf(os.Stderr, "无法启动服务器: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("用法: goai <命令> [参数]")
	fmt.Println("命令:")
	fmt.Println("  generate [--model <模型名>] \"你的提示\"   - 生成一次性内容")
	fmt.Println("  chat [--model <模型名>]                  - 启动交互式聊天")
	fmt.Println("  server [--port <端口号>]                 - 启动API服务器")
}
