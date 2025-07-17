# GitHub Actions 工作流

此目录包含 GoAI 项目的 GitHub Actions 工作流。

## CI/CD 流水线 (`cicd.yml`)

该工作流能自动化构建 Go 应用程序、创建 Docker 镜像并将其推送到 Docker Hub 的整个过程。

### 工作原理

- **触发器:** 当您在仓库中创建一个格式为 `v*.*.*` (例如 `v1.0.0`, `v1.2.3`) 的新 Git 标签时，该工作流将自动触发。
- **作业 (Jobs):**
    1.  **`build-and-push`**:
        - 从代码仓库中检出最新的代码。
        - 设置 Go 环境 (版本 1.22)。
        - 使用存储在 GitHub Secrets 中的凭据登录到 Docker Hub。
        - 使用项目根目录下的 `Dockerfile` 文件构建 Docker 镜像。
        - 将镜像推送到 Docker Hub，并附带两个标签：
            - `latest`: 始终指向 `main` 分支的最新构建。
            - `git-sha`: 与 Git 提交哈希对应的唯一标签，用于版本追溯。

### 设置说明

要使用此工作流，您需要执行以下一次性设置：

1.  **更新 `cicd.yml` 文件:**
    - 打开 `.github/workflows/cicd.yml` 文件。
    - 在 `Build and push Docker image` 步骤中，将两处 `DOCKERHUB_USERNAME` 替换为您自己的 Docker Hub 用户名。

    ```yaml
    # ...
          tags: |
            你的DOCKERHUB用户名/goai:latest
            你的DOCKERHUB用户名/goai:${{ github.sha }}
    ```

2.  **配置 GitHub Secrets:**
    - 进入您项目的 GitHub 仓库页面。
    - 导航到 `Settings` > `Secrets and variables` > `Actions`。
    - 点击 `New repository secret` 添加以下两个 secret：
        - **`DOCKERHUB_USERNAME`**: 您的 Docker Hub 用户名。
        - **`DOCKERHUB_TOKEN`**: 您的 Docker Hub 访问令牌 (Access Token)。您可以从 Docker Hub 账户的 "Security" 设置页面生成。

完成这些步骤后，每当您向 `main` 分支推送代码时，该工作流都将自动运行。

---

## Release 构建 (`release.yml`)

该工作流用于在创建新的版本标签时，自动为多个平台构建可执行文件，并创建一个包含这些文件的 GitHub Release。

### 工作原理

- **触发器:** 当您在仓库中创建一个格式为 `v*.*.*` (例如 `v1.0.0`, `v1.2.3`) 的新 Git 标签时，该工作流将自动触发。

- **作业 (Jobs):**
    1.  **`build-and-release`**:
        - 使用构建矩阵 (Build Matrix) 并行地为以下三个目标平台进行构建：
            - Windows (amd64)
            - Linux (amd64)
            - macOS (amd64)
        - 从 `cmd/cli/main.go` 构建统一的可执行文件 (`goai.exe` 或 `goai`)。
        - 自动创建一个与 Git 标签同名的 GitHub Release。
        - 将所有平台的可执行文件作为附件 (Assets) 上传到这个 Release 中。

### 使用方法

1.  在您的本地仓库中，使用以下命令创建一个新的版本标签：
    ```bash
    git tag v1.0.0
    ```
    (请将 `v1.0.0` 替换为您想要的实际版本号)

2.  将这个标签推送到 GitHub：
    ```bash
    git push origin v1.0.0
    ```

推送成功后，GitHub Actions 将自动开始执行此工作流。您可以在仓库的 "Actions" 标签页查看进度，并在 "Releases" 页面找到生成的最终文件。

### 设置说明

此工作流无需额外配置。它使用 `GITHUB_TOKEN` 这个由 GitHub 自动提供的 secret 来授权创建 Release 和上传附件。
