package main

import (
	"EnGin/internal/db"
	"EnGin/internal/global"
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "执行数据库 Schema 自动迁移",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.InitDB(global.Config.Database.Driver, global.Config.Database.Dsn())
		if err != nil {
			global.Log.Fatal("连接数据库失败", zap.Error(err))
		}
		defer db.Close()
		global.Log.Info("开始执行数据库 Schema 自动迁移...")

		ctx := context.Background()
		// 直接使用 db.Client 跨包调用
		if err := db.Client.Schema.Create(ctx); err != nil {
			global.Log.Fatal("Schema 迁移失败", zap.Error(err))
		}

		global.Log.Info("数据库 Schema 迁移执行完毕！")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
