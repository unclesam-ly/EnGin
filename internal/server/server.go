package server

import (
	"EnGin/internal/global"
	"EnGin/internal/handler"
	"EnGin/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() {
	gin.SetMode(global.Config.Server.Mode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.TraceID())
	r.Use(middleware.Logger())

	// 注册路由
	handler.RegisterRoutes(r)
	port := global.Config.Server.Port
	global.Log.Info(fmt.Sprintf("Server is running on port %d", port))
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		global.Log.Fatal("Web 服务启动失败", zap.Error(err))
	}
}
