package main

import (
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
		err := db.Init(global.Config.Database.Driver, global.Config.Database.Dsn())
		if err != nil {
			global.Log.Fatal("数据库连接失败", zap.Error(err))
		}
		defer db.Close()

		global.Log.Info("数据库初始化成功，正在启动 Web 服务...")
		// 启动 Gin 服务
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
