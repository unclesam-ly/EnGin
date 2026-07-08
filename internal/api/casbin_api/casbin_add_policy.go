package casbin_api

import (
	"EnGin/internal/db"
	"EnGin/internal/global"
	"EnGin/internal/middleware"
	"EnGin/internal/utils/res"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (CasbinApi) CasbinAddPolicyView(c *gin.Context) {
	cr := middleware.GetBind[PolicyRequest](c)

	isAdded, err := db.GlobalEnforcer.AddPolicy(cr.RoleID, cr.Path, cr.Method)
	if err != nil {
		global.Log.Error("添加策略失败", zap.Error(err))
		res.InternalError("添加策略失败", c)
		return
	}

	if !isAdded {
		res.OkWithMsg("策略已存在", c)
		return
	}

	res.OkWithMsg("策略添加成功", c)
}
