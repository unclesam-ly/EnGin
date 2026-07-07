package middleware

import (
	"EnGin/internal/db"
	"EnGin/internal/ent/user"
	"EnGin/internal/global"
	"EnGin/internal/service/redis_ser"
	"EnGin/internal/utils/jwts"
	"EnGin/internal/utils/res"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Auth(c *gin.Context) {
	accessToken := c.GetHeader("access_token")
	claims, err := jwts.CheckAccessToken(accessToken)
	if err != nil {
		res.Unauthorized("认证失败", c)
		c.Abort()
		return
	}

	if redis_ser.HasLogout(accessToken) {
		res.FailWithMsg("当前登录已经注销", c)
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}

func AuthAdmin(c *gin.Context) {
	accessToken := c.GetHeader("access_token")
	if accessToken == "" {
		res.Unauthorized("认证失败", c)
		c.Abort()
		return
	}

	claims, err := jwts.CheckAccessToken(accessToken)
	if err != nil {
		res.Unauthorized("认证失败或 Token 已过期", c)
		c.Abort()
		return
	}

	// 预加载用户的关联角色列表 (WithRoles)
	ctx := c.Request.Context()
	u, err := db.Client.User.Query().
		Where(user.ID(int(claims.UserID))).
		WithRoles(). // 预加载关系
		Only(ctx)

	if err != nil {
		res.Forbidden("用户不存在", c)
		c.Abort()
		return
	}

	// 遍历用户的角色列表，检查是否有 "admin" 角色代码
	isAdmin := false
	for _, role := range u.Edges.Roles {
		if role.Code == "admin" {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		res.Forbidden("需要管理员权限", c)
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}

func GetAuth(c *gin.Context) (cl *jwts.MyClaims) {
	cl = new(jwts.MyClaims)
	_claims, ok := c.Get("claims")
	if !ok {
		return nil
	}
	cl, ok = _claims.(*jwts.MyClaims)
	return
}

func JwtRefresh(c *gin.Context) {
	token := c.Request.Header.Get("refresh_token")
	if token == "" {
		res.BadRequest("未携带refresh_token", c)
		c.Abort()
		return
	}

	// 解析token
	claims, err := jwts.CheckRefreshToken(token)
	if err != nil {
		res.FailWithMsg("refresh_token 无效", c)
		c.Abort()
		return
	}

	// 检查redis
	if redis_ser.HasLogout(token) {
		res.FailWithMsg("当前登录已经注销", c)
		c.Abort()
		return
	}

	c.Set("refresh_claims", claims)
	c.Next()
}

func CasbinMiddleware(c *gin.Context) {
	// 获取用户信息
	claims := GetAuth(c)
	if claims == nil {
		res.Forbidden("权限不足", c)
		c.Abort()
		return
	}

	// 获取请求路径方法
	obj := c.Request.URL.Path
	act := c.Request.Method

	// 使用 user_ID 作为主体进行权限检查
	sub := fmt.Sprintf("user_%d", claims.UserID)

	// 查看用户的角色
	roles, _ := db.GlobalEnforcer.GetRolesForUser(sub)

	// 检查权限
	success, err := db.GlobalEnforcer.Enforce(sub, obj, act)
	if err != nil {
		global.Log.Error("权限检查出错", zap.Error(err))
		res.InternalError("权限检查出错", c)
		c.Abort()
		return
	}

	if !success {
		global.Log.Warn("权限不足",
			zap.String("sub", sub),
			zap.String("obj", obj),
			zap.String("act", act),
			zap.Strings("roles", roles))
		res.Forbidden("权限不足", c)
		c.Abort()
		return
	}

	c.Next()
}
