package db

import (
	"EnGin/internal/global"
	"fmt"

	"github.com/casbin/casbin/v3"
	entadapter "github.com/casbin/ent-adapter"
)

// GlobalEnforcer 全局的 Casbin 权限强制执行器
var GlobalEnforcer *casbin.Enforcer

// InitCasbin 初始化 Casbin 并绑定数据库适配器
func InitCasbin() error {
	// 初始化 ent-adapter，它会自动使用您配置的数据库并在库中生成 casbin_rule 表
	adapter, err := entadapter.NewAdapter(global.Config.Database.Driver, global.Config.Database.Dsn())
	if err != nil {
		return fmt.Errorf("初始化 Casbin 数据库适配器失败: %w", err)
	}

	// 加载本地的配置文件和数据库适配器
	enforcer, err := casbin.NewEnforcer("configs/casbin.conf", adapter)
	if err != nil {
		return fmt.Errorf("创建 Casbin Enforcer 失败: %w", err)
	}

	// 从数据库加载权限规则
	if err := enforcer.LoadPolicy(); err != nil {
		return fmt.Errorf("加载 Casbin 策略失败: %w", err)
	}

	GlobalEnforcer = enforcer
	global.Log.Info("Casbin 初始化成功")

	return nil
}
