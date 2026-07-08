package casbin_api

import (
	"EnGin/internal/db"
	"EnGin/internal/global"
	"EnGin/internal/middleware"
	"EnGin/internal/utils/res"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (CasbinApi) CasbinRemoveRoleInheritanceView(c *gin.Context) {
	cr := middleware.GetBind[RoleInheritanceRequest](c)

	// 移除角色继承关系: g,roleID,inheritFrom
	isRemoved, err := db.GlobalEnforcer.RemoveGroupingPolicy(cr.RoleID, cr.InheritFrom)
	if err != nil {
		global.Log.Error("移除角色继承关系失败", zap.Error(err))
		res.InternalError("移除角色继承关系失败", c)
		return
	}

	if !isRemoved {
		res.NotFound("角色关系继承不存在", c)
		return
	}

	res.OkWithMsg("角色继承关系移除成功", c)
}
