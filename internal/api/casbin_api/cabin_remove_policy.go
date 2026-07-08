package casbin_api

import (
	"EnGin/internal/db"
	"EnGin/internal/global"
	"EnGin/internal/middleware"
	"EnGin/internal/utils/res"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (CasbinApi) CasbinRemovePolicyView(c *gin.Context) {
	cr := middleware.GetBind[PolicyRequest](c)

	isRemoved, err := db.GlobalEnforcer.RemovePolicy(cr.RoleID, cr.Path, cr.Method)
	if err != nil {
		global.Log.Error("删除策略失败", zap.Error(err))
		res.InternalError("删除策略失败", c)
		return
	}

	if !isRemoved {
		res.NotFound("策略不存在", c)
		return
	}

	res.OkWithMsg("策略删除成功", c)
}
