
## TODO
- 实现`--template "<template>"`: 使用 Go 模板语法格式化最终的提示。模板中可通过 `{{.prompt}}` 和 `{{.text}}` 引用输入。

## 更新日志

### v0.2.2 （20250805）
- 启动server服务，主页给出API使用说明。

### v0.2.2
- 重构代码，更为灵活的调用各类AI。

### v0.1.3
- fix: API请求时默认使用stream模式，以防过长文本返回结果不完整。

### v0.1.2
- 完善release.yml

### v0.1.1
- 测试研究github action

### v0.1.0
- 初版本发布