package casbin_api

import (
	"EnGin/internal/db"
	"EnGin/internal/global"
	"EnGin/internal/utils/res"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (CasbinApi) CasbinRoleInheritanceListView(c *gin.Context) {
	// 获取所以角色关系
	policies, err := db.GlobalEnforcer.GetGroupingPolicy()
	if err != nil {
		global.Log.Error("查询角色继承关系失败", zap.Error(err))
		res.InternalError("查询角色继承关系失败", c)
		return
	}

	if policies == nil {
		policies = [][]string{}
	}

	res.OkWithData(policies, c)
}
