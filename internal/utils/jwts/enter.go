package jwts

import (
	"EnGin/internal/global"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Roles    []uint `json:"roles"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type RefreshClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type MyClaims struct {
	Claims
	jwt.RegisteredClaims
}

// GenAccessToken 生成Access Token
func GenAccessToken(data Claims) (string, error) {
	claims := MyClaims{
		Claims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.AccessExpire) * time.Hour)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(global.Config.Jwt.AccessSecret))
}

// GenRefreshToken 生成Refresh Token
func GenRefreshToken(userID uint) (string, error) {
	claims := RefreshClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.RefreshExpire) * time.Hour)),
			Issuer:    global.Config.Jwt.Issuer + "_refresh",
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(global.Config.Jwt.RefreshSecret))
}

// CheckAccessToken 验证访问令牌
func CheckAccessToken(token string) (*MyClaims, error) {
	claims, err := parseToken(token, global.Config.Jwt.AccessSecret, &MyClaims{})
	if err != nil {
		return nil, err
	}

	return claims.(*MyClaims), nil
}

// CheckRefreshToken 验证刷新令牌
func CheckRefreshToken(token string) (*RefreshClaims, error) {
	claims, err := parseToken(token, global.Config.Jwt.RefreshSecret, &RefreshClaims{})
	if err != nil {
		return nil, err
	}

	return claims.(*RefreshClaims), nil
}

// parseToken 通用token解析函数
func parseToken(token, secret string, claims jwt.Claims) (jwt.Claims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if tokenObj.Valid {
		return tokenObj.Claims, nil
	}

	return nil, errors.New("token无效")
}
