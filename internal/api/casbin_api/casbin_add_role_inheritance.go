package casbin_api

import (
	"EnGin/internal/db"
	"EnGin/internal/global"
	"EnGin/internal/middleware"
	"EnGin/internal/utils/res"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (CasbinApi) CasbinAddRoleInheritanceView(c *gin.Context) {
	cr := middleware.GetBind[RoleInheritanceRequest](c)

	// 添加角色继承关系
	isAdded, err := db.GlobalEnforcer.AddGroupingPolicy(cr.RoleID, cr.InheritFrom)
	if err != nil {
		global.Log.Error("添加角色继承关系失败", zap.Error(err))
		res.InternalError("添加角色继承关系失败", c)
		return
	}

	if !isAdded {
		res.OkWithMsg("角色继承关系已经存在", c)
		return
	}

	res.OkWithMsg("角色继承关系添加成功", c)
}
