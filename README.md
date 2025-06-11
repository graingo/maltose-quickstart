# Maltose 快速开始模板

基于 Maltose 框架的标准项目结构模板

## 项目结构

```
maltose-quickstart/
├── api/            # API 定义和请求/响应结构
│   └── v1/         # API 版本
├── cmd/            # 应用程序入口
│   ├── server.go   # 服务启动
│   ├── config.go   # (可选) 配置加载及钩子(Hook)函数
│   └── route.go    # 路由注册
├── config/         # 配置文件
├── internal/       # 内部代码，不对外暴露
│   ├── controller/ # 控制器，处理请求和响应
│   ├── service/    # 服务接口定义
│   ├── logic/      # 业务逻辑实现
│   ├── dao/        # 数据访问对象，处理数据库操作
│   ├── provider/   # 封装对第三方服务的调用（如微服务、SaaS API）
│   └── model/      # 模型定义
│       ├── entity/ # 实体定义，映射数据库表结构
│       └── do/     # (可选) 数据对象，用于承载复杂查询结果
├── utility/        # 工具函数
└── main.go         # 主入口文件
```

## 代码分层

本项目遵循清晰的分层架构，各层职责如下：

### 表现层（API 和 Controller）

- **API 层 (`api/`)**: 定义请求和响应的数据结构，包含路由元数据
- **控制器层 (`internal/controller/`)**: 处理 HTTP 请求，参数验证，调用 Service 层，返回响应

### 业务层（Service）

- **服务层 (`internal/service/`)**: 定义业务逻辑的接口 (`interface`)。
- **逻辑层 (`internal/logic/`)**: 业务逻辑的具体实现层。它实现了 `Service` 层定义的接口，并负责编排 `DAO` 和 `Provider` 来完成业务。

### 数据与外部服务层（DAO, Model, Provider）

- **DAO 层 (`internal/dao/`)**: 封装对本项目数据库的访问逻辑。
- **Provider 层 (`internal/provider/`)**: 封装对所有外部服务（其他微服务、第三方 API）的客户端调用。
- **模型层 (`internal/model/`)**:
  - **实体 (`entity/`)**: 与数据库表结构一一对应，由工具生成，不应手动修改。
  - **数据对象 (`do/`)**: (可选) 手动创建，用于 `DAO` 层承载自定义的复杂查询结果。

### 基础设施层

- **入口 (`cmd/`)**: 负责应用的启动、路由注册和配置加载。
- **工具 (`utility/`)**: 通用工具函数和辅助功能
- **配置 (`config/`)**: 应用配置文件

## 数据流

1. **请求** → **Router** → **Controller**
2. **Controller** → **Service/Logic**
3. **Service/Logic** → **DAO** (访问数据库) 或 **Provider** (调用外部服务)
4. **DAO** → **数据库**
5. **Provider** → **外部 API(微服务等)**
6. **响应** 按相反方向返回

## 开发指南

1. 定义 API 数据结构 (`api/v1/*.go`)
2. 在 `internal/service/` 中定义业务接口
3. 如需调用外部服务，在 `internal/provider/` 中实现封装
4. 实现 Controller 处理请求 (`internal/controller/*.go`)
5. 在 `internal/logic/` 中实现业务逻辑，按需调用 DAO 或 Provider
6. 实现 DAO 处理数据访问 (`internal/dao/*.go`)
7. 定义 Entity 和 (可选的) DO 数据模型 (`internal/model/*/*.go`)
8. 注册路由 (`cmd/route.go`)

## 启动项目

```bash
go run main.go
```
