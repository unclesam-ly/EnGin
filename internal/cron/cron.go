package cron

import (
	"EnGin/internal/global"
	"time"

	"github.com/robfig/cron/v3"
)

// Cron 实例全局持有，方便后续管理（比如优雅退出时关闭）

var Crontab *cron.Cron

func InitCron() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	Crontab = cron.New(cron.WithSeconds(), cron.WithLocation(timezone))

	// 在这里注册您的定时任务
	// Crontab.AddFunc("*/5 * * * * *", func() {
	//     global.Log.Info("定时任务执行中...")
	// })

	Crontab.Start()
	global.Log.Info("后台定时任务启动成功")
}

func Stop() {
	if Crontab != nil {
		Crontab.Stop()
		global.Log.Info("后台定时任务已安全停止")
	}
}
