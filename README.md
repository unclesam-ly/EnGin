# EnGin

基于 **entgo + Gin + Cobra** 的 Go 后端项目脚手架，开箱即用。

## 技术栈

| 组件 | 选型 |
|------|------|
| ORM | [entgo.io/ent](https://entgo.io) — 代码生成式 ORM |
| Web 框架 | [Gin](https://github.com/gin-gonic/gin) |
| CLI | [Cobra](https://github.com/spf13/cobra) |
| 配置管理 | [Viper](https://github.com/spf13/viper) + YAML + 热加载 |
| 日志 | [Zap](https://github.com/uber-go/zap) + Lumberjack 自动轮转 |
| 数据库 | MySQL / PostgreSQL |

## 项目结构

```
EnGin/
├── cmd/server/              # CLI 入口与子命令
│   ├── main.go              # 程序入口
│   ├── root.go              # 根命令（加载配置、初始化日志）
│   ├── server.go            # server 子命令：启动 Web 服务
│   ├── migrate.go           # migrate 子命令：数据库迁移
│   ├── user.go              # user 子命令：用户操作
│   └── gen.go               # gen 子命令：重新生成 ent 代码
├── configs/
│   └── config.yaml          # 主配置文件
├── internal/
│   ├── conf/                # 配置结构定义与加载
│   ├── db/                  # 数据库客户端初始化 + 连接池
│   ├── ent/                 # ent 生成的 ORM 代码
│   │   └── schema/          # Schema 定义（在这里定义数据模型）
│   ├── global/              # 全局变量 (Config, Log)
│   ├── handler/             # HTTP 路由注册
│   ├── logger/              # Zap 日志初始化
│   ├── server/              # Gin 服务启动
│   └── service/             # 业务逻辑层
├── Makefile
├── go.mod
└── .air.toml                # Air 热重载配置
```

## 快速开始

### 环境要求

- Go >= 1.25
- MySQL 或 PostgreSQL

### 1. 配置数据库

编辑 `configs/config.yaml`：

```yaml
database:
  driver: "postgres"     # mysql 或 postgres
  host: "127.0.0.1"
  port: 5432
  user: "postgres"
  password: "123456"
  dbname: "test_db"
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 运行

```bash
# 启动 Web 服务
go run cmd/server/*.go server

# 或使用 Makefile
make run
```

## CLI 命令

```bash
# 查看帮助
go run cmd/server/*.go -h

# 启动 HTTP 服务（默认 :8080）
go run cmd/server/*.go server -c configs/config.yaml

# 数据库自动迁移
go run cmd/server/*.go migrate -c configs/config.yaml

# 重新生成 ent 代码
go run cmd/server/*.go gen

# 用户操作
go run cmd/server/*.go user create    # 创建用户
```

## 开发指南

### 添加新数据模型

1. 在 `internal/ent/schema/` 下创建新的 Schema 文件
2. 运行代码生成：

```bash
make generate
# 或 go generate ./internal/ent
```

### 添加新 API

1. 在 `internal/service/` 中编写业务逻辑
2. 在 `internal/handler/router.go` 中注册路由

### 配置说明

| 配置项 | 说明 |
|--------|------|
| `server.port` | 服务端口 |
| `server.mode` | `debug`（控制台彩色日志）/ `release`（JSON 日志） |
| `database.driver` | `mysql` 或 `postgres` |
| `logger.level` | `debug` / `info` / `warn` / `error` |
| `logger.filename` | 日志输出路径 |
| `logger.max_size` | 单个日志文件最大大小（MB） |
| `logger.max_backups` | 保留的旧日志文件数 |
| `logger.max_age` | 日志保留天数 |

### 热重载（可选）

项目已配置 [Air](https://github.com/air-verse/air)，安装后可直接使用：

```bash
air
```

## Makefile

```makefile
make run        # 启动服务
make generate   # 生成 ent 代码
make build      # 编译二进制到 bin/server
```

## License

MIT
