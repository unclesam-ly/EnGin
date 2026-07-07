package router

import (
	"EnGin/internal/api"
	"EnGin/internal/api/user_api"
	"EnGin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	app := api.App.UserApi
	g.GET("users",
		middleware.BindQueryMiddleware[user_api.UserListRequest],
		middleware.AuthAdmin,
		app.UserListView)
}
