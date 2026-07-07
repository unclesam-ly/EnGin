package main

import (
	"bufio"
	"context"
	"ent-scaffold/internal/db"
	"ent-scaffold/internal/global"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "后台用户管理工具",
}

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "交互式创建一个管理员用户",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Init(global.Config.Database.Driver, global.Config.Database.Dsn())
		if err != nil {
			global.Log.Fatal("连接数据库失败", zap.Error(err))
		}
		defer db.Close()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请输入用户名: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		fmt.Print("请输入密码: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)
		if username == "" || password == "" {
			fmt.Println("用户名和密码不能为空")
			return
		}

		ctx := context.Background()
		// 直接使用 db.Client 跨包调用并创建数据
		u, err := db.Client.User.Create().
			// 根据您的 schema 字段做适当的设定，如：
			// SetUsername(username).
			// SetPassword(password).
			Save(ctx)
		if err != nil {
			global.Log.Error("创建用户失败", zap.Error(err))
			fmt.Println("用户创建失败: ", err.Error())
			return
		}

		fmt.Printf("✨ 用户创建成功！ID: %d, 用户名: %s\n", u.ID, username)
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)
	rootCmd.AddCommand(userCmd)
}
