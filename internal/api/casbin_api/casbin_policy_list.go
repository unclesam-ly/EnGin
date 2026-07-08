package casbin_api

import (
	"EnGin/internal/db"
	"EnGin/internal/global"
	"EnGin/internal/utils/res"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PolicyListResponse struct {
	Policies         [][]string `json:"policies"`
	GroupingPolicies [][]string `json:"grouping_policies"`
}

func (CasbinApi) CasbinPolicyListView(c *gin.Context) {
	// 获取所有策略
	policies, err := db.GlobalEnforcer.GetPolicy()
	if err != nil {
		global.Log.Error("获取策略失败", zap.Error(err))
		res.FailWithMsg("获取策略失败", c)
		return
	}

	if policies == nil {
		policies = [][]string{}
	}

	// 获取所有角色继承关系
	groupingPolicies, err := db.GlobalEnforcer.GetGroupingPolicy()
	if err != nil {
		global.Log.Error("获取继承策略失败", zap.Error(err))
		res.FailWithMsg("获取继承策略失败", c)
		return
	}

	if groupingPolicies == nil {
		groupingPolicies = [][]string{}
	}

	policyAllList := PolicyListResponse{
		Policies:         policies,
		GroupingPolicies: groupingPolicies,
	}

	res.OkWithData(policyAllList, c)
}
