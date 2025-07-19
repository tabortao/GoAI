# docker部署goai教程

> GoAI 是一个多功能的AI文本处理服务，通过单一代码库同时提供标准的HTTP API和功能丰富的命令行（CLI）两种交互模式。项目核心基于 Go 语言和 LangChainGo 框架，支持通过 `config.json` 动态配置和切换多种大型语言模型（LLM），并通过 Docker 实现一键式容器化部署。

## Docker部署

喜欢使用docker-compose来部署docker，新建[docker-compose.yml](https://github.com/tabortao/GoAI/blob/main/docker-compose.yml)如下。

```bash
services:
  goai:
    image: dk.nastool.de/tabortoa/goai #国内用户镜像加速 dk.nastool.de/tabortoa/goai
    restart: always
    ports:
      - "1388:8080" # 端口号1388可自己修改
    command: ["server"]
    volumes:
      - ./config.json:/root/config.json 
```
下载[config.json.example](https://github.com/tabortao/GoAI/blob/main/config.json.example)，修改为config.json，然后填入自己的AI Token信息。

### 阿里云服务器部署

SSH登录阿里云服务器，输入下列命令
```bash
cd opt/opt/docker
mkdir goai
# 上传docker-compose.yml和congig.json文件到goai文件夹
docker-compose up -d
# 如需更新，请执行下方命令
# docker-compose pull && docker-compose up -d 
```

### 飞牛部署
- 打开飞牛Docker-Compose-新建项目。
- 项目名称：输入goai
- 路径:选择自己存放dockr的文件夹，新建文件夹goai，把文件docker-compose.yml、config.json放到goai文件夹。
- 选择创建项目后立即启动，点击启动。
