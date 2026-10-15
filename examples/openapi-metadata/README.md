# 文档元数据示例

此示例使用 Maltose v0.5.0 的 `Extensions.Info/Server/Tag`。模板依赖与 CLI 配套使用 v0.5.1，再按以下步骤接入。

在项目根目录创建 `apidocs/configure.go`：

```go
package apidocs

import "github.com/graingo/maltose/net/mhttp/contract"

func Configure(e *contract.Extensions) error {
    if err := e.Info(contract.Info{
        Title: "Maltose Quickstart API",
        Version: "0.1.0",
        Description: "Maltose 快速开始示例接口",
        Contact: &contract.Contact{Name: "API Team"},
    }); err != nil { return err }
    if err := e.Server(contract.Server{URL: "/", Description: "当前服务"}); err != nil { return err }
    return e.Tag(contract.Tag{Name: "公共服务", Description: "问候与告别示例"})
}
```

使用项目的实际 module 路径指定扩展包：

```bash
maltose gen openapi -s api -o cmd/openapi.yaml --openapi-version 3.1.0 \
  --extensions github.com/graingo/maltose-quickstart/apidocs
maltose gen openapi -s api -o cmd/openapi.yaml --openapi-version 3.1.0 \
  --extensions github.com/graingo/maltose-quickstart/apidocs --check
go test -race ./...
```

产物包含以下元数据；完整文档和清单均由生成器写入：

```yaml
info:
  title: Maltose Quickstart API
  version: 0.1.0
  description: Maltose 快速开始示例接口
  contact:
    name: API Team
servers:
  - url: /
    description: 当前服务
tags:
  - name: 公共服务
    description: 问候与告别示例
```

现有 Server 继续 embed 同一对产物，无需修改加载代码。元数据变更会改变 manifest 文档摘要，接口指纹保持一致。

只需标题和版本时，可以省略扩展包并使用 `--title "Maltose Quickstart API" --api-version 0.1.0`。CLI 参数与扩展声明的相同字段必须一致；空字段最后采用默认标题 `API` 和版本 `1.0.0`。`--api-version` 表示接口文档版本，`--openapi-version` 表示 OpenAPI 规范版本。

生成与 `--check` 始终使用相同参数。已配置元数据的项目修改 module 路径后，也应带上新的扩展包路径重新生成。直接调用 Go `contract.Generate` 时通过配置回调声明元数据，旧 `Options.Title` 迁移为 `Extensions.Info`。
