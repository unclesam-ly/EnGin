package router

import (
	"EnGin/internal/api"
	"EnGin/internal/api/casbin_api"
	"EnGin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CasbinRouter(g *gin.RouterGroup) {
	app := api.App.CasbinApi
	g.GET("policy", middleware.AuthAdmin, app.CasbinPolicyListView)

	g.POST("policy",
		middleware.AuthAdmin,
		middleware.BindJsonMiddleware[casbin_api.PolicyRequest],
		app.CasbinAddPolicyView)

	g.DELETE("policy",
		middleware.AuthAdmin,
		middleware.BindJsonMiddleware[casbin_api.PolicyRequest],
		app.CasbinRemovePolicyView)

	// 角色继承权限路由
	g.POST("role_inheritance",
		middleware.AuthAdmin,
		middleware.BindJsonMiddleware[casbin_api.RoleInheritanceRequest],
		app.CasbinAddRoleInheritanceView)

	g.DELETE("role_inheritance",
		middleware.AuthAdmin,
		middleware.BindJsonMiddleware[casbin_api.RoleInheritanceRequest],
		app.CasbinRemoveRoleInheritanceView)

	g.GET("role_inheritance", middleware.AuthAdmin, app.CasbinRoleInheritanceListView)
}
