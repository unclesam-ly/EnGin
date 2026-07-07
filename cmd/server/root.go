package main

import (
	"EnGin/internal/conf"
	"EnGin/internal/global"
	"EnGin/internal/logger"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "EnGin",
	Short: "EnGin 命令行脚手架服务",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
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
