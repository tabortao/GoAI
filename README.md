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

GoAI 提供了两种使用方式：命令行工具和 HTTP API。

### 命令行工具 (CLI)

所有命令都通过 `goai` 可执行文件运行。

#### `server` - 启动 API 服务

启动 HTTP API 服务器。
```bash
./goai server
```
服务将在 `.env` 文件中配置的 `HTTP_PORT` 上运行。

#### `generate` - 生成文本

从单个提示生成文本。
```bash
./goai generate "写一首关于宇宙的诗"
```
**参数:**
- `--stream`: 启用流式输出，实时返回内容。
- `--model <name>`: 指定要使用的模型 (如 `openai` 或 `ollama`)。

**示例:**
```bash
# 使用 Ollama 模型进行流式生成
./goai generate "你好，世界！" --model ollama --stream
```

#### `chat` - 交互式聊天

启动一个交互式聊天会话。
```bash
./goai chat
```
**参数:**
- `--model <name>`: 指定聊天会话中使用的模型。

在会话中，输入 `exit` 或 `quit` 来结束。

### HTTP API

首先，请确保 API 服务正在运行 (`./goai server`)。

#### `POST /api/v1/generate`

接收文本生成请求。

**请求体 (JSON):**
```json
{
  "prompt": "用户的输入文本",
  "stream": false,
  "model": "openai"
}
```
- `prompt` (string, required): 输入的提示文本。
- `stream` (boolean, optional): 是否以流式返回，默认为 `false`。
- `model` (string, optional): 指定模型，如果留空则使用默认配置的模型。

**示例 (cURL):**

- **非流式请求:**
  ```bash
  curl -X POST http://localhost:8080/api/v1/generate \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "解释一下什么是黑洞",
    "model": "openai"
  }'
  ```

- **流式请求:**
  ```bash
  curl -X POST http://localhost:8080/api/v1/generate \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "给我讲一个关于旅行的短故事",
    "stream": true
  }'
  ```
  响应将是 `text/event-stream` 格式。

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/generate" -Method Post -Headers @{"Content-Type"="application/json"} -Body '{"prompt":"tell me a story","stream":true}'
```


#### `GET /health`

健康检查端点，用于监控服务状态。
```bash
curl http://localhost:8080/health
```
**响应:**
```json
{
  "status": "ok"
}
```

## 🐳 Docker 部署

使用 Docker Compose 可以轻松启动 GoAI 服务及其依赖（如 Ollama）。

1.  **确保 `.env` 文件已配置。**

2.  **构建并启动服务:**
    ```bash
    docker-compose up --build
    ```
    此命令将：
    - 构建 `goai` 服务的 Docker 镜像。
    - 启动 `goai` 容器并运行 API 服务。
    - （可选）启动一个 `ollama` 服务容器。

3.  **在后台运行:**
    ```bash
    docker-compose up -d
    ```

4.  **停止服务:**
    ```bash
    docker-compose down
    ```

## 🛠️ 技术栈

- **语言**: Go 1.21+
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