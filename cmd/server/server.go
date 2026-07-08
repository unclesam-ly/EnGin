package main

import (
	"EnGin/internal/cron"
	"EnGin/internal/db"
	"EnGin/internal/global"
	"EnGin/internal/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动 Web API 接口服务",
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化 db 包
		err := db.InitDB(global.Config.Database.Driver, global.Config.Database.Dsn())
		if err != nil {
			global.Log.Fatal("数据库连接失败", zap.Error(err))
		}
		defer db.Close()

		// 初始化 Redis
		err = db.InitRedis(global.Config.Redis.Addr, global.Config.Redis.Password, global.Config.Redis.Db)
		if err != nil {
			global.Log.Fatal("初始化 Redis 失败", zap.Error(err))
		}
		defer db.CloseRedis()

		// 初始化 Casbin（会自动连接数据库并在数据库中建立 casbin_rules 表）
		err = db.InitCasbin()
		if err != nil {
			global.Log.Fatal("初始化 Casbin 失败", zap.Error(err))
		}

		// 初始化定时任务
		cron.InitCron()
		defer cron.Stop()

		// 启动 Gin 服务
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
