package middleware

import (
	"EnGin/internal/db"
	"EnGin/internal/utils/res"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LimitConfig 限流配置
type LimitConfig struct {
	MaxRequest int           // 最大请求数
	TimeWindow time.Duration // 时间窗口
	Message    string        // 提示信息
}

var (
	DefaultLimit = LimitConfig{
		MaxRequest: 100,
		TimeWindow: 60 * time.Second,
		Message:    "请求过于频繁，请稍后重试",
	}

	SearchLimit = LimitConfig{
		MaxRequest: 20,
		TimeWindow: 60 * time.Second,
		Message:    "搜索过于频繁，请稍后重试",
	}

	AuthLimit = LimitConfig{
		MaxRequest: 5,
		TimeWindow: 60 * time.Second,
		Message:    "登录尝试过多，请稍后重试",
	}

	PaymentLimit = LimitConfig{
		MaxRequest: 3,
		TimeWindow: 60 * time.Second,
		Message:    "操作过于频繁，请稍后重试",
	}
)

// RateLimitMiddleware ip限流器
func RateLimitMiddleware(c *gin.Context, config LimitConfig) {
	ip := c.ClientIP()
	key := fmt.Sprintf("rate_limit:%s", ip)
	ctx := context.Background()

	// 获取当前计数
	count, err := db.Redis.Get(ctx, key).Int()
	if err != nil && err.Error() != "redis: nil" {
		// Redis 错误，放行请求（降级策略）
		c.Next()
		return
	}

	// 检查是否超限
	if count >= config.MaxRequest {
		res.FailWithMsg("请求过于频繁，请稍后重试", c)
		c.Abort()
		return
	}

	// 增加计数
	pipe := db.Redis.Pipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, config.TimeWindow)
	_, err = pipe.Exec(ctx)

	if err != nil {
		// Redis 错误，放行请求
		c.Next()
		return
	}

	c.Next()
}

// 快捷方法

func DefaultRateLimit(c *gin.Context) {
	RateLimitMiddleware(c, DefaultLimit)
}

func SearchRateLimit(c *gin.Context) {
	RateLimitMiddleware(c, SearchLimit)
}

func AuthRateLimit(c *gin.Context) {
	RateLimitMiddleware(c, AuthLimit)
}

func PaymentRateLimit(c *gin.Context) {
	RateLimitMiddleware(c, AuthLimit)
}
