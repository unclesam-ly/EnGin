package main

import (
	"EnGin/internal/db"
	"EnGin/internal/ent"
	"EnGin/internal/ent/role"
	"EnGin/internal/global"
	"EnGin/internal/utils/pwd"
	"bufio"
	"context"
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
		hashedPwd, err := pwd.GenerateFromPassword(password)
		if err != nil {
			fmt.Errorf("密码加密错误 %w", err)
			return
		}
		if username == "" || password == "" {
			fmt.Println("用户名和密码不能为空")
			return
		}

		ctx := context.Background()

		// 1. 尝试查询 admin 角色
		adminRole, err := db.Client.Role.Query().
			Where(role.CodeEQ("admin")).
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				// 2. 如果不存在 admin 角色，自动创建
				adminRole, err = db.Client.Role.Create().
					SetCode("admin").
					SetName("超级管理员").
					Save(ctx)
				if err != nil {
					global.Log.Error("创建管理员角色失败", zap.Error(err))
					fmt.Println("创建管理员角色失败: ", err.Error())
					return
				}
				fmt.Println("👉 系统中未发现 admin 角色，已自动创建。")
			} else {
				global.Log.Error("查询角色失败", zap.Error(err))
				fmt.Println("查询角色失败: ", err.Error())
				return
			}
		}

		// 3. 直接使用 db.Client 创建用户并挂载角色
		u, err := db.Client.User.Create().
			SetUsername(username).
			SetPassword(hashedPwd).
			AddRoles(adminRole). // 💡 重点：挂载角色
			Save(ctx)
		if err != nil {
			global.Log.Error("创建用户失败", zap.Error(err))
			fmt.Println("用户创建失败: ", err.Error())
			return
		}

		fmt.Printf("✨ 用户创建成功！ID: %d, 用户名: %s, 角色: admin\n", u.ID, username)
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)
	rootCmd.AddCommand(userCmd)
}
