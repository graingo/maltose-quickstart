# Maltose 快速开始模板

基于 Maltose 框架的标准项目结构模板

## 项目结构

```
maltose-quickstart/
├── api/            # API 定义和请求/响应结构
│   └── v1/         # API 版本
├── cmd/            # 应用程序入口
│   ├── server.go   # 服务器启动
│   └── config.go   # 配置加载
├── config/         # 配置文件
├── internal/       # 内部代码，不对外暴露
│   ├── controller/ # 控制器，处理请求和响应
│   ├── service/    # 服务层，实现业务逻辑
│   ├── dao/        # 数据访问对象，处理数据库操作
│   └── model/      # 模型定义
│       ├── entity/ # 实体定义，映射数据库表结构
│       └── do/     # 数据对象，用于业务数据传输
├── route/          # 路由定义
│   └── route.go    # 路由注册
├── utility/        # 工具函数
└── main.go         # 主入口文件
```

## 代码分层

本项目遵循清晰的分层架构，各层职责如下：

### 表现层（API 和 Controller）

- **API 层 (`api/`)**: 定义请求和响应的数据结构，包含路由元数据
- **控制器层 (`internal/controller/`)**: 处理 HTTP 请求，参数验证，调用 Service 层，返回响应

### 业务层（Service）

- **服务层 (`internal/service/`)**: 实现业务逻辑，协调多个 DAO 操作，不关心 HTTP 细节

### 数据层（DAO 和 Model）

- **DAO 层 (`internal/dao/`)**: 封装数据访问逻辑，包括数据库操作
- **模型层 (`internal/model/`)**:
  - **实体 (`entity/`)**: 与数据库表结构一一对应
  - **数据对象 (`do/`)**: 用于业务数据传输和扩展处理

### 基础设施层

- **路由 (`route/`)**: 注册路由及中间件
- **工具 (`utility/`)**: 通用工具函数和辅助功能
- **配置 (`config/`)**: 应用配置文件

## 数据流

1. 请求 → Router → Controller
2. Controller → Service (使用 DO 对象传递数据)
3. Service → DAO (使用 DO 和 Entity 对象)
4. DAO → 数据库 (使用 Entity 对象)
5. 响应按相反方向返回

## 开发指南

1. 定义 API 数据结构和路由元数据 (`api/v1/*.go`)
2. 实现 Controller 处理请求 (`internal/controller/*.go`)
3. 实现 Service 处理业务逻辑 (`internal/service/*.go`)
4. 实现 DAO 处理数据访问 (`internal/dao/*.go`)
5. 定义 Entity 和 DO 数据模型 (`internal/model/*/*.go`)
6. 注册路由 (`route/route.go`)

## 启动项目

```bash
go run main.go
```
