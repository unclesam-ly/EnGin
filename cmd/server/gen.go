package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// findProjectRoot 从当前工作目录向上查找包含 go.mod 的项目根目录
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err != nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("未找到 go.mod 所在的项目根目录")
		}
		dir = parent
	}
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成 Ent ORM 的模板代码",
	// 覆盖父命令的 PersistentPreRunE，生成代码不需要加载 config.yaml
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		rootDir, err := findProjectRoot()
		if err != nil {
			println("定位项目根目录失败: " + err.Error())
			return
		}
		println("正在运行 go generate ./internal/ent...")
		c := exec.Command("go", "generate", "./internal/ent")
		c.Dir = rootDir // 强制在项目根目录下执行
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
