package api

import (
	"GoAI/internal/config"
	"GoAI/internal/llm"
	"context"
	"fmt"
	
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
)

// GenerateRequest 定义了 /api/v1/generate 请求的结构
type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt" binding:"required"`
	Text   string `json:"text"`
	Stream bool   `json:"stream"` // 默认为 false，但我们的处理将优先考虑流式
}

// GenerateHandler 处理生成内容的请求
func GenerateHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GenerateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 默认开启流式响应
		stream := true

		modelConfig, err := cfg.GetModelConfig(req.Model)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("获取模型配置失败: %v", err)})
			return
		}

		llmManager, err := llm.NewManager(modelConfig)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建LLM管理器"})
			return
		}

		finalPrompt := req.Prompt
		if req.Text != "" {
			finalPrompt = fmt.Sprintf("%s\n\n---\n\n%s", req.Prompt, req.Text)
		}

		messages := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, "你是一个专业的AI助手。请根据用户的指令完成任务。"),
			llms.TextParts(llms.ChatMessageTypeHuman, finalPrompt),
		}

		if stream {
			c.Writer.Header().Set("Content-Type", "text/event-stream")
			c.Writer.Header().Set("Cache-Control", "no-cache")
			c.Writer.Header().Set("Connection", "keep-alive")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

			// 使用回调函数处理流式数据
			callback := func(ctx context.Context, chunk []byte) error {
				if _, err := c.Writer.Write(chunk); err != nil {
					// 客户端可能已经断开连接
					return err
				}
				c.Writer.Flush()
				return nil
			}

			_, err := llmManager.GenerateContent(context.Background(), messages, modelConfig.Temperature, 4096, callback)
			if err != nil {
				// 由于响应头已发送，我们不能再发送JSON错误，但可以在日志中记录
				fmt.Fprintf(c.Writer, "Error: %v\n", err)
			}
		} else {
			// 非流式响应
			response, err := llmManager.GenerateContent(context.Background(), messages, modelConfig.Temperature, 4096, nil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("生成内容失败: %v", err)})
				return
			}
			c.JSON(http.StatusOK, gin.H{"response": response})
		}
	}
}