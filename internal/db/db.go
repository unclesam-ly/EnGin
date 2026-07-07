package db

import (
	"database/sql"
	"fmt"
	"time"

	"ent-scaffold/internal/ent"

	entsql "entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
	_ "github.com/lib/pq"              // PostgreSQL 驱动
)

// Client 跨包调用的统一数据库客户端实例
var Client *ent.Client

// Init 根据配置初始化数据库连接
func Init(driver string, dsn string) error {
	// 先创建标准 sql.DB 以配置连接池
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return fmt.Errorf("打开数据库连接失败 [%s]: %w", driver, err)
	}

	// 连接池设置
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	// 将 sql.DB 包装为 ent 客户端
	drv := entsql.OpenDB(driver, db)
	Client = ent.NewClient(ent.Driver(drv))
	return nil
}

// Close 关闭数据库连接
func Close() {
	if Client != nil {
		Client.Close()
	}
}
