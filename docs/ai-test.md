
## 运行环境测试
Invoke-RestMethod -Uri "https://go.xxx.top/api/v1/generate" `
  -Method Post `
  -ContentType "application/json; charset=utf-8" `
  -Body (@{
    "prompt" = "请将以下文本翻译成英语，不要有过多的描述："
    "text" = "GoAI 是一个多功能的AI文本处理服务，通过单一代码库同时提供标准的HTTP API和功能丰富的命令行（CLI）两种交互模式。项目核心基于 Go 语言和 LangChainGo 框架，支持通过 `config.json` 动态配置和切换多种大型语言模型（LLM），并通过 Docker 实现一键式容器化部署。"
    "model" = "goai-chat"
  } | ConvertTo-Json -Compress)

## 开发环境测试
### 不设置模型
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/generate" `
  -Method Post `
  -ContentType "application/json; charset=utf-8" `
  -Body (@{
    "prompt" = "请将以下文本翻译成英语，不要有过多的描述："
    "text" = "GoAI 是一个多功能的AI文本处理服务，通过单一代码库同时提供标准的HTTP API和功能丰富的命令行（CLI）两种交互模式。项目核心基于 Go 语言和 LangChainGo 框架，支持通过 `config.json` 动态配置和切换多种大型语言模型（LLM），并通过 Docker 实现一键式容器化部署。"
  } | ConvertTo-Json -Compress)

### 自定义模型
Invoke-RestMethod -Uri "http://192.168.3.4:1388/api/v1/generate" `
  -Method Post `
  -ContentType "application/json; charset=utf-8" `
  -Body (@{
    "prompt" = "请将以下文本翻译成英语，不要有过多的描述："
    "text" = "GoAI 是一个多功能的AI文本处理服务，通过单一代码库同时提供标准的HTTP API和功能丰富的命令行（CLI）两种交互模式。项目核心基于 Go 语言和 LangChainGo 框架，支持通过 `config.json` 动态配置和切换多种大型语言模型（LLM），并通过 Docker 实现一键式容器化部署。"
    "model" = "deepseek-chat"
  } | ConvertTo-Json -Compress)
