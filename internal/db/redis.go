package db

import (
	"EnGin/internal/global"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Redis 统一导出的 Redis 客户端实例
var Redis *redis.Client

// InitRedis 初始化 Redis 连接
func InitRedis(addr string, password string, dbIndex int) error {
	if addr == "" {
		return nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbIndex,
	})
	// 立即 Ping 一下，验证连接是否成功
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("连接 Redis 失败: %w", err)
	}

	Redis = rdb
	global.Log.Info("Redis 初始化成功")

	return nil
}

// CloseRedis 关闭 Redis 连接
func CloseRedis() {
	if Redis != nil {
		Redis.Close()
	}
}
