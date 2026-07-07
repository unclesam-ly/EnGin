package server

import (
	"EnGin/internal/global"
	"EnGin/internal/middleware"
	"EnGin/internal/router"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() {
	gin.SetMode(global.Config.Server.Mode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.TraceID())
	r.Use(middleware.Logger())

	// 注册路由
	router.RegisterRoutes(r)
	port := global.Config.Server.Port
	global.Log.Info(fmt.Sprintf("服务运行在端口 %d", port))
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		global.Log.Fatal("Web 服务启动失败", zap.Error(err))
	}
}
