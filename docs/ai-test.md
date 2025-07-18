
## 运行环境测试
Invoke-RestMethod -Uri "https://go.sdgarden.top/api/v1/generate" `
  -Method Post `
  -ContentType "application/json; charset=utf-8" `
  -Body (@{
    "prompt" = "请将以下文本翻译成英语，不要有过多的描述："
    "text" = "人工智能是当今科技领域最热门的话题之一"
  } | ConvertTo-Json -Compress)

## 开发环境测试
### 不设置模型
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/generate" `
  -Method Post `
  -ContentType "application/json; charset=utf-8" `
  -Body (@{
    "prompt" = "请将以下文本翻译成英语，不要有过多的描述："
    "text" = "人工智能是当今科技领域最热门的话题之一"
  } | ConvertTo-Json -Compress)

### 自定义模型
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/generate" `
  -Method Post `
  -ContentType "application/json; charset=utf-8" `
  -Body (@{
    "prompt" = "请将以下文本翻译成英语，不要有过多的描述："
    "text" = "人工智能是当今科技领域最热门的话题之一"
    "provider" = "deepseek_official"
    "model" = "deepseek-chat"
  } | ConvertTo-Json -Compress)