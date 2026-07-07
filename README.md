# EnGin

[English](#english) | [简体中文](#简体中文) | [日本語](#日本語)

---

## English

A production-ready Go backend project scaffolding based on **Ent ORM + Gin + Cobra**, out of the box.

### Features

* **ORM**: [Ent](https://entgo.io/) - An entity framework for Go. Strongly-typed, code-generated schema definitions,
  automatic migration, and clean graph queries.
* **Router**: [Gin](https://github.com/gin-gonic/gin) - High-performance HTTP routing.
* **CLI Command**: [Cobra](https://github.com/spf13/cobra) - Rich CLI subcommands (`server`, `migrate`, `user`, `gen`).
* **Configurations**: [Viper](https://github.com/spf13/viper) - YAML configuration with hot-reloading support.
* **Database**: Supported MySQL and PostgreSQL.
* **Cache / Token Blacklist**: Redis integration (go-redis/v9) for logout blacklisting.
* **Authentication**: Dual-token (Access Token & Refresh Token) authentication.
* **Type-Safe Binding**: Generic HTTP request binding & auto-validation middleware.
* **Structured Logs**: Zap logger + Lumberjack log rotation with project-level prefixes.

### Directory Structure

```
EnGin/
├── cmd/server/              # CLI Entry and commands
│   ├── main.go              # App entry point
│   ├── root.go              # Root command (initializes config & logs)
│   ├── server.go            # server command: Starts Web server
│   ├── migrate.go           # migrate command: DB auto-migrations
│   ├── user.go              # user command: CRUD & setup admin user
│   └── gen.go               # gen command: Runs ent generate
├── configs/
│   └── config.yaml          # Main configuration file
├── internal/
│   ├── api/                 # API controllers (grouped by domain)
│   │   ├── auth_api/        # Authentication API (login, token refresh)
│   │   └── user_api/        # User management API
│   ├── conf/                # Viper config mapping definitions
│   ├── db/                  # DB connection pool (Ent & Redis initialization)
│   ├── ent/                 # Ent generated ORM files
│   │   └── schema/          # Ent schema definitions (models)
│   ├── global/              # Global constants and shared variables
│   ├── middleware/          # Gin middlewares (CORS, Logger, JWT, generic Bind)
│   ├── logger/              # Zap logger configuration
│   ├── router/              # Gin route group registrations
│   └── service/             # Domain business logic layer
├── Makefile                 # Quick command shortcuts
├── go.mod                   # Dependency definitions
└── .air.toml                # Air hot-reloading configuration
```

### Getting Started

#### Prerequisites

* Go >= 1.25
* MySQL or PostgreSQL
* Redis (Optional, needed for token blacklist verification)

#### 1. Setup Configuration

Edit `configs/config.yaml` to fill in database details:

```yaml
database:
    driver: "postgres"     # postgres or mysql
    host: "127.0.0.1"
    port: 5432
    user: "postgres"
    password: "your_password"
    dbname: "engin_db"
```

#### 2. Install Dependencies

```bash
go mod tidy
```

#### 3. Run Commands

```bash
# Run migrations (creates tables and intermediate join tables)
go run cmd/server/*.go migrate

# Create an administrator interactively (automatically binds admin role)
go run cmd/server/*.go user create

# Start HTTP server
go run cmd/server/*.go server
```

---

## 简体中文

基于 **Ent ORM + Gin + Cobra** 的开箱即用型 Go 后端项目脚手架。

### 项目特性

* **ORM**：[Ent](https://entgo.io/) - 强类型、代码生成式的图结构实体框架，支持智能迁移与极简联表。
* **Web 框架**：[Gin](https://github.com/gin-gonic/gin) - 高性能 HTTP 路由分发。
* **命令行**：[Cobra](https://github.com/spf13/cobra) - 规范 of CLI 命令工具集（`server` 运行、`migrate` 迁移、`user` 管理、
  `gen` 生成）。
* **配置管理**：[Viper](https://github.com/spf13/viper) - 支持 YAML 映射与动态热加载。
* **缓存与注销**：集成 Redis (go-redis/v9)，支持基于黑名单的 Token 安全注销。
* **安全鉴权**：支持双 Token (Access Token & Refresh Token) 安全刷新机制。
* **强类型绑定**：内置基于 Go 泛型（Generics）的请求绑定与自动参数校验中间件。
* **工程日志**：基于 Zap + Lumberjack 自动滚动归档，且支持按配置自动添加项目日志前缀。

### 目录结构

```
EnGin/
├── cmd/server/              # CLI 命令行入口与子命令
│   ├── main.go              # 应用程序入口
│   ├── root.go              # 根命令（加载配置与初始化日志）
│   ├── server.go            # server 命令：运行 Web 接口服务
│   ├── migrate.go           # migrate 命令：自动执行表结构迁移
│   ├── user.go              # user 命令：管理用户/交互式创建管理员
│   └── gen.go               # gen 命令：触发 ent 静态代码生成
├── configs/
│   └── config.yaml          # 全局主配置文件
├── internal/
│   ├── api/                 # API 控制器层（按业务模块划分）
│   │   ├── auth_api/        # 鉴权相关接口 (登录、刷新Token)
│   │   └── user_api/        # 用户管理接口
│   ├── conf/                # 配置映射结构体定义
│   ├── db/                  # 数据库连接池（Ent 和 Redis 连接初始化）
│   ├── ent/                 # Ent 框架生成的底层 CRUD 代码
│   │   └── schema/          # Schema 模型定义（数据表结构声明处）
│   ├── global/              # 全局通用变量与常量定义
│   ├── middleware/          # Gin 中间件库 (跨域、日志拦截、JWT、泛型参数绑定)
│   ├── logger/              # Zap 日志引擎配置
│   ├── router/              # Gin 路由树注册与子路由分发
│   └── service/             # 核心领域业务逻辑层
├── Makefile                 # 快捷开发指令集
├── go.mod                   # 模块依赖描述
└── .air.toml                # Air 热重载（自动编译）配置文件
```

### 快速开始

#### 环境要求

* Go >= 1.25
* MySQL 或 PostgreSQL
* Redis (可选，用于黑名单过滤)

#### 1. 配置参数

编辑 `configs/config.yaml` 填写对应数据库信息：

```yaml
database:
    driver: "postgres"     # mysql 或 postgres
    host: "127.0.0.1"
    port: 5432
    user: "postgres"
    password: "你的密码"
    dbname: "engin_db"
```

#### 2. 下载依赖

```bash
go mod tidy
```

#### 3. 运行指南

```bash
# 自动生成数据库表（包括 M2M 中间表）
go run cmd/server/*.go migrate

# 交互式创建一个管理员用户（自动赋予 admin 角色）
go run cmd/server/*.go user create

# 启动 Web 服务
go run cmd/server/*.go server
```

---

## 日本語

**Ent ORM + Gin + Cobra** に基づく、すぐに使えるプロダクション対応の Go バックエンドプロジェクトスケルトン（脚手架）です。

### 主な機能と特徴

* **ORM**: [Ent](https://entgo.io/) - 強力な静的型付けコード生成型のグラフ構造 ORM。簡単なリレーション構築とマイグレーションをサポート。
* **Web フレームワーク**: [Gin](https://github.com/gin-gonic/gin) - 高速な HTTP ルーティング。
* **CLI コマンド**: [Cobra](https://github.com/spf13/cobra) - 構造化されたサブコマンド体系（`server`, `migrate`, `user`,
  `gen`）。
* **設定管理**: [Viper](https://github.com/spf13/viper) - YAML 設定ファイルのマッピングとホットリロードのサポート。
* **キャッシュ / ログアウト**: Redis (go-redis/v9) によるログアウト Token ブラックリスト管理。
* **認証**: デュアルトークン (Access Token & Refresh Token) による安全な認証スキーム。
* **型安全バインディング**: Go のジェネリクス（Generics）を活用したリクエスト解析＆自動バリデーションミドルウェア。
* **ログ記録**: Zap + Lumberjack ログローテーション、プロジェクト固有のログプレフィックス設定対応。

### ディレクトリ構成

```
EnGin/
├── cmd/server/              # CLI コマンドの起点とサブコマンド定義
│   ├── main.go              # アプリケーションのメインエントリー
│   ├── root.go              # ルートコマンド（設定読み込み・ログ初期化）
│   ├── server.go            # server コマンド: Web サービスの起動
│   ├── migrate.go           # migrate コマンド: データベースのマイグレーション
│   ├── user.go              # user コマンド: 管理者ユーザーの作成・管理
│   └── gen.go               # gen コマンド: ent コードの自動生成実行
├── configs/
│   └── config.yaml          # グローバル設定ファイル
├── internal/
│   ├── api/                 # API コントローラ層（ビジネスドメイン別）
│   │   ├── auth_api/        # 認証関連 API (ログイン、トークン更新)
│   │   └── user_api/        # ユーザー管理 API
│   ├── conf/                # 設定ファイルの構造体マッピング定義
│   ├── db/                  # データベース接続プール（Ent と Redis の初期化）
│   ├── ent/                 # Ent 生成コード
│   │   └── schema/          # テーブルモデル定義（スキーマ定義場所）
│   ├── global/              # グローバル定数と共通変数
│   ├── middleware/          # Gin ミドルウェア (CORS, ログ, JWT, ジェネリクスバインディング)
│   ├── logger/              # Zap ロガーの設定
│   ├── router/              # ルートグループ登録とルーティング定義
│   └── service/             # サービス層（ビジネスロジック実行）
├── Makefile                 # クイックコマンド定義
├── go.mod                   # モジュール依存関係定義
└── .air.toml                # Air ホットリロード設定ファイル
```

### クイックスタート

#### 動作要件

* Go >= 1.25
* MySQL または PostgreSQL
* Redis (オプション、ログアウトブラックリスト機能に必要)

#### 1. 設定ファイルの編集

`configs/config.yaml` のデータベース接続情報を設定します：

```yaml
database:
    driver: "postgres"     # mysql または postgres
    host: "127.0.0.1"
    port: 5432
    user: "postgres"
    password: "your_password"
    dbname: "engin_db"
```

#### 2. 依存関係のインストール

```bash
go mod tidy
```

#### 3. コマンドの実行

```bash
# マイグレーションを実行し、中間テーブルなどを自動生成
go run cmd/server/*.go migrate

# 対話形式で管理者アカウントを作成（自動で admin ロールを付与）
go run cmd/server/*.go user create

# Web サービスを起動
go run cmd/server/*.go server
```
