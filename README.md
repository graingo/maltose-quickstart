# Maltose Quickstart

`maltose-quickstart` 是 Maltose 的官方项目模板，提供一个无需外部服务、克隆后即可运行的 HTTP API 示例。

## 环境要求

- Go 1.25 或更高版本
- Maltose CLI（通过 CLI 创建项目时需要）

```bash
go install github.com/graingo/maltose/cmd/maltose@latest
```

## 创建项目

推荐通过 CLI 创建项目，并显式设置正式 Module Path：

```bash
maltose new hello-maltose --module github.com/acme/hello-maltose
cd hello-maltose
go run .
```

也可以直接克隆本仓库：

```bash
git clone https://github.com/graingo/maltose-quickstart.git
cd maltose-quickstart
go run .
```

服务默认监听 `http://127.0.0.1:8081`。

## 调用示例

```bash
curl 'http://127.0.0.1:8081/api/v1/hello?name=Maltose'
```

响应使用 Maltose 标准响应结构：

```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "name": "Hello, Maltose!",
    "tags": []
  }
}
```

模板提供以下入口：

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/health` | 健康检查 |
| GET | `/api/v1/hello?name=Maltose` | 查询参数绑定与业务调用 |
| POST | `/api/v1/bye` | JSON 请求绑定 |
| POST | `/api/v2/bye` | API 版本路由 |
| GET | `/swagger` | Swagger UI |
| GET | `/api.json` | OpenAPI 文档 |

POST 请求示例：

```bash
curl -X POST 'http://127.0.0.1:8081/api/v2/bye' \
  -H 'Content-Type: application/json' \
  -d '{"name":"Maltose"}'
```

## 项目结构

```text
.
├── api/hello/              # API 契约，按业务模块与版本组织
│   ├── v1/
│   └── v2/
├── cmd/                    # Server 装配与路由注册
├── config/config.yaml      # 应用配置
├── internal/
│   ├── controller/         # HTTP 请求适配
│   ├── logic/              # Service 接口实现与业务编排
│   ├── model/              # 内部输入输出模型
│   ├── service/            # 业务能力接口
│   ├── dao/                # 可选：数据库访问
│   └── provider/           # 可选：外部服务封装
├── utility/                # 可选：项目通用工具
└── main.go                 # 应用生命周期入口
```

请求处理流程为：

```text
Router → Controller → Service/Logic → DAO 或 Provider
```

`main.go` 使用 `m.App` 管理 HTTP Server 的生命周期，并注入同一个 Logger 记录启动和关闭过程。业务项目增加数据库、Redis、文件 Logger 或可观测性 Provider 后，应通过 `m.WithCloser` 注册这些资源。

## 开发与验证

```bash
go test -race ./...
go vet ./...
go build ./...
```

修改 `api/` 下的接口定义后，可以使用 Maltose CLI 生成对应骨架：

```bash
maltose gen service
```

模板依赖固定在 `go.mod` 中。升级 Maltose 后应执行完整验证：

```bash
go get github.com/graingo/maltose@latest
go mod tidy
go test -race ./...
```

完整文档见 [Maltose Documentation](https://graingo.github.io/maltose-docs/)。

## API 契约与文档

控制器的路径、状态、包装和字段来源由 `m.Meta` 与 Req/Res 声明。示例显式使用 `envelope:"maltose"`，查询字段使用 `query`。

```bash
maltose gen openapi -s api -o cmd/openapi.yaml --openapi-version 3.1.0
maltose gen openapi -s api -o cmd/openapi.yaml --check
```

提交文档与 `cmd/openapi.yaml.manifest.json`。服务通过 embed 加载二者，启动时比较实际契约，`/api.json` 和 Swagger 使用加载的文档。修改模块路径后重新生成；`maltose new` 会自动完成这一步。

模板使用 Maltose v0.5.1 的契约编译器，配套安装同版本 CLI：

```bash
go install github.com/graingo/maltose/cmd/maltose@v0.5.1
```

从 v0.3 升级时，先按 [API 契约迁移说明](https://graingo.github.io/maltose-docs/components/server/api-metadata.html) 调整 DTO，再重新生成并校验文档。

v0.5.0 的文档元数据能力及接入步骤见 [文档元数据示例](examples/openapi-metadata/README.md)，包含标题、API 版本、联系人、服务器地址和标签说明。框架与 CLI 使用相同版本。

`HelloRes.Tags` 使用 `json:"tags" field:"required"`：字段必有、非 null，允许空数组。控制器返回 `[]string{}`，保证空结果编码为 `[]`。需要 null 的字段显式声明 `field:"required,nullable"`。v0.5.1 更新了响应契约指纹，升级后重新生成文档及 manifest；生成器采用应用依赖的框架版本，框架与 CLI 一起升级。
