package redis_ser

import (
	"EnGin/internal/db"
	"EnGin/internal/global"
	"EnGin/internal/utils/jwts"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// Logout 标记当前令牌已注销 (拉入黑名单)
func Logout(accessToken, refreshToken string) {
	atoken, err := jwts.CheckAccessToken(accessToken)
	rtoken, err := jwts.CheckRefreshToken(refreshToken)
	if err != nil {
		global.Log.Error("token验证失败", zap.Error(err))
		return
	}
	// 存储 access_token 黑名单 key
	accessKey := fmt.Sprintf("logout_access_%s", accessToken)
	// 存储 refresh_token 黑名单 key
	refreshKey := fmt.Sprintf("logout_refresh_%s", refreshToken)
	// 计算剩余有效期并作为 Redis 的过期时间缓存进去
	db.Redis.Set(context.Background(), accessKey, "", atoken.ExpiresAt.Sub(time.Now()))
	db.Redis.Set(context.Background(), refreshKey, "", rtoken.ExpiresAt.Sub(time.Now()))
}

// HasLogout 检查该 token 是否已经注销
func HasLogout(accessToken string) bool {
	// 如果 Redis 没有开启，默认放行
	if db.Redis == nil {
		return false
	}

	accessKey := fmt.Sprintf("logout_access_%s", accessToken)
	_, err := db.Redis.Get(context.Background(), accessKey).Result()
	if err == nil {
		return true // 如果能查到 key，说明已经被注销
	}

	return false
}
