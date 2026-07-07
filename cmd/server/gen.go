package main

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成 Ent ORM 的模板代码",
	Run: func(cmd *cobra.Command, args []string) {
		println("正在运行 go generate ./internal/ent...")

		c := exec.Command("go", "generate", "./internal/ent")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			println("代码生成失败: " + err.Error())
			return
		}

		println("Ent ORM 代码生成成功！")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
