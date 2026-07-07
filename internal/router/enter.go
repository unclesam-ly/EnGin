package router

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 主路由入口，注册中间件并分发子路由
func RegisterRoutes(r *gin.Engine) {
	// 分组并调用子路由
	g := r.Group("api/v1")
	AuthRouter(g)
	UserRouter(g)
}
