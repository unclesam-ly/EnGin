package router

import (
	"EnGin/internal/api"
	"EnGin/internal/api/auth_api"
	"EnGin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(g *gin.RouterGroup) {
	app := api.App.AuthApi

	g.POST("login",
		middleware.BindJsonMiddleware[auth_api.LoginRequest],
		middleware.AuthRateLimit,
		app.LoginView)
}
