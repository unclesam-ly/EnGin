package main

import (
	"EnGin/internal/conf"
	"EnGin/internal/global"
	"EnGin/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "EnGin",
	Short: "EnGin 命令行脚手架服务",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 如果在子目录（如 cmd/server）启动，自动切换工作目录到 go.mod 所在的项目根目录
		if rootDir, err := findProjectRoot(); err == nil {
			_ = os.Chdir(rootDir)
		}

		cfg, err := conf.LoadConfig(cfgFile)
		if err != nil {
			return err
		}

		global.Config = cfg
		log, err := logger.Init(cfg.Logger)
		if err != nil {
			return err
		}

		global.Log = log
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "configs/config.yaml", "指定 YAML 配置文件路径")
}
