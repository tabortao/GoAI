# GoAI - AI 文本处理服务

GoAI 是一个多功能的AI文本处理服务，通过单一代码库同时提供标准的HTTP API和功能丰富的命令行（CLI）两种交互模式。项目核心基于 Go 语言和 LangChainGo 框架，支持动态切换多种大型语言模型（LLM），并通过 Docker 实现一键式容器化部署。

## ✨ 项目特点

- **双模式支持**: 同时提供 RESTful API 服务和交互式命令行工具，共享核心业务逻辑。
- **高度灵活性**: 支持通过环境变量动态配置 LLM 提供商（如 OpenAI, Ollama），轻松切换模型。
- **生产就绪**: 内置结构化日志、统一配置管理、API 健康检查和优雅关机机制。
- **可扩展架构**: 采用模块化和接口驱动的设计，易于维护和功能扩展。
- **容器化部署**: 提供 `Dockerfile` 和 `docker-compose.yml`，实现一键式构建和部署。

## 🚀 快速开始

### 1. 环境准备

- [Go](https://golang.org/dl/) (版本 1.21 或更高)
- [Docker](https://www.docker.com/get-started) 和 [Docker Compose](https://docs.docker.com/compose/install/) (用于容器化部署)

### 2. 安装与配置

1.  **克隆项目**
    ```bash
    git clone https://github.com/your-username/GoAI.git
    cd GoAI
    ```

2.  **配置环境变量**
    复制环境变量示例文件，并根据你的需求进行修改。
    ```bash
    cp .env.example .env
    ```
    编辑 `.env` 文件，至少配置一个 LLM 提供商：
    ```dotenv
    # OpenAI API Key
    OPENAI_API_KEY=sk-your_openai_api_key_here

    # Ollama URL (如果使用)
    OLLAMA_URL=http://localhost:11434

    # 服务端口
    HTTP_PORT=8080

    # 日志级别 (DEBUG, INFO, WARN, ERROR)
    LOG_LEVEL=INFO
    ```

3.  **安装依赖**
    ```bash
    go mod tidy
    ```

### 3. 构建可执行文件

```bash
# 在 Windows 上
go build -o goai.exe ./cmd/cli

# 在 Linux 或 macOS 上
go build -o goai ./cmd/cli
```

## 🔧 使用说明

GoAI 支持三种主要的使用方式：**命令行 (CLI)**、**HTTP API** 和 **Docker 部署**。下面将详细介绍每种方式的使用方法和示例。

### 1. 命令行工具 (CLI)

所有命令都通过 `goai` 可执行文件运行。请确保您已根据“快速开始”部分的指引完成了编译。

#### 启动 API 服务

```bash
.\goai.exe server
```
此命令会启动一个 HTTP 服务器，监听在 `.env` 文件中配置的 `HTTP_PORT` 端口（默认为 `8080`）。

#### 文本生成 (`generate`)

这是最核心的命令，用于执行各种文本生成任务。

**基础用法:**
```bash
.\goai.exe generate "写一首关于宇宙的诗"
```

**常用参数:**
- `--model <name>`: 指定要使用的模型 (如 `openai` 或 `ollama`)。
- `--stream`: 启用流式输出，实时显示结果。
- `--text <content>`: 提供额外的长文本内容，通常与 `prompt` 结合使用。
- `--template "<template>"`: 使用 Go 模板语法格式化最终的提示。模板中可通过 `{{.prompt}}` 和 `{{.text}}` 引用输入。

**场景示例:**

*   **翻译任务:**
    ```bash
    .\goai.exe generate "请将以下文本翻译成英语：" --text "你好，我的祖国是中国！"
    ```

*   **文本摘要 (流式输出):**
    ```bash
    .\goai.exe generate "请为以下文章生成摘要：" --text "人工智能（AI）是研究、开发用于模拟、延伸和扩展人的智能的理论、方法、技术及应用系统的一门新的技术科学...（此处省略长文本）" --stream
    ```

*   **使用模板进行角色扮演:**
    ```bash
    .\goai.exe generate "你是一位经验丰富的软件工程师。" --text "请解释一下什么是 RESTful API？" --template "角色设定: {{.prompt}}\n\n问题: {{.text}}"
    ```

#### 交互式聊天 (`chat`)

启动一个可以持续对话的交互式会话。

```bash
.\goai.exe chat --model goai-chat
```
在会话中，输入 `exit` 来结束。

### 2. HTTP API

请先使用 `.\goai.exe server` 或 Docker 启动服务。API 提供了与 CLI 类似的功能。

#### 端点: `POST /api/v1/generate`

**请求体 (JSON):**
```json
{
  "prompt": "用户的输入文本",
  "text": "可选的额外文本",
  "template": "可选的模板字符串，例如：指令: {{.prompt}}",
  "stream": false,
  "model": "goai-chat"
}
```

**场景示例:**

*   **翻译任务 (cURL):**
    ```bash
    curl -X POST http://localhost:8080/api/v1/generate \
    -H "Content-Type: application/json" \
    -d '{
      "prompt": "请将以下文本翻译成英语：",
      "text": "你好，我的祖国是中国！",
      "model": "openai"
    }'
    ```

*   **翻译任务 (PowerShell):**
    ```powershell
    Invoke-RestMethod -Uri "http://localhost:8080/api/v1/generate" `
      -Method Post `
      -ContentType "application/json; charset=utf-8" `
      -Body (@{
        "prompt" = "请将以下文本翻译成英语："
        "text" = "你好，我的祖国是中国！"
        "model" = "openai"
      } | ConvertTo-Json -Compress)
    ```

*   **文本摘要 (cURL, 流式):**
    ```bash
    curl -X POST http://localhost:8080/api/v1/generate \
    -H "Content-Type: application/json" \
    -d '{
      "prompt": "请为以下文章生成摘要：",
      "text": "人工智能是当今科技领域最热门的话题之一...（此处省略长文本）",
      "stream": true
    }'
    ```

*   **使用模板 (PowerShell):**
    ```powershell
    Invoke-RestMethod -Uri "http://localhost:8080/api/v1/generate" `
      -Method Post `
      -ContentType "application/json; charset=utf-8" `
      -Body (@{
        "prompt" = "你是一位资深美食家。"
        "text" = "请推荐一道适合夏天的菜。"
        "template" = "角色: {{.prompt}}\n\n任务: {{.text}}"
      } | ConvertTo-Json -Compress)
    ```

### 3. Docker 部署与使用

使用 Docker Compose 是推荐的生产环境部署方式。

1.  **启动服务:**
    确保 `.env` 文件已根据您的需求配置好。
    ```bash
    docker-compose up --build -d
    ```
    `-d` 参数使服务在后台运行。

2.  **通过 API 与服务交互:**
    Docker 服务启动后，API 会暴露在您主机的 `8080` 端口上。您可以像在本地一样，使用 `curl` 或 `Invoke-RestMethod` 等工具直接调用 `http://localhost:8080`。

    **示例 (从您的主机直接调用 Docker 内的服务):**
    ```bash
    curl -X POST http://localhost:8080/api/v1/generate \
    -H "Content-Type: application/json" \
    -d '{"prompt": "Docker 容器内运行的服务，你好！"}'
    ```

3.  **在 Docker 容器内执行 CLI 命令:**
    您也可以进入正在运行的 `goai` 容器，直接执行 CLI 命令。

    ```bash
    # docker exec -it <容器名称或ID> <命令>
    docker exec -it goai-app ./goai generate "在 Docker 容器内向我问好"
    ```
    *注意: `goai-app` 是 `docker-compose.yml` 中定义的服务名，可能会因您的设置而异。*

4.  **查看日志:**
    ```bash
    docker-compose logs -f goai-app
    ```

5.  **停止服务:**
    ```bash
    docker-compose down
    ```

## 🛠️ 技术栈

- **语言**: Go 1.24+
- **核心框架**: 
  - **HTTP 服务**: Gin
  - **命令行**: Cobra
  - **LLM 集成**: LangChainGo
- **配置管理**: Viper
- **日志**: slog (Go 原生结构化日志库)
- **部署**: Docker, Docker Compose

## 📂 项目结构

```
GoAI/
├── cmd/
│   ├── cli/
│   │   └── main.go        # 命令行应用入口
│   └── server/
│       └── main.go        # HTTP服务独立入口
├── internal/
│   ├── api/               # API 定义 (handler, router)
│   ├── cli/               # Cobra 命令定义
│   ├── config/            # 配置加载
│   ├── core/              # 核心业务逻辑
│   ├── llm/               # LLM 客户端管理器
│   └── models/            # 数据结构定义
├── pkg/
│   └── utils/             # 通用工具 (logger, template)
├── .env.example           # 环境变量示例
├── Dockerfile             # 生产镜像构建文件
├── docker-compose.yml     # Docker Compose 编排文件
├── go.mod
└── README.md              # 项目说明
```