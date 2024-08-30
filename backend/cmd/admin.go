package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
	"strings"
	"wordma/server/model"
)

func NewCreateAdminCommand() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:     "admin",
		Aliases: []string{"admin"},
		Short:   "创建管理员账号",
		Long:    Banner,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			// 输入管理员用户名
			fmt.Print("请输入管理员用户名: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			// 输入管理员邮箱
			fmt.Print("请输入管理员邮箱: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			// 输入密码
			fmt.Print("请输入密码: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			// 输入重复密码
			fmt.Print("请重复密码: ")
			confirmPassword, _ := reader.ReadString('\n')
			confirmPassword = strings.TrimSpace(confirmPassword)

			// 简单的密码匹配检查
			if password != confirmPassword {
				fmt.Println("错误：两次输入的密码不匹配。")
				return
			}
			userData, err := model.GetUserByEmail(email)
			if userData != nil {
				fmt.Println("邮箱已存在")
				return
			}
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println("数据库查询错误：", err)
				return
			}
			// 创建管理员用户
			user := &model.User{
				Name:     username,
				Email:    email,
				Password: password,
			}
			err = user.SetPasswordEncrypt(password)
			if err != nil {
				fmt.Println("密码加密失败：", err)
				return
			}
			err = model.CreateUser(user)
			if err != nil {
				fmt.Println("创建用户失败：", err)
				return
			}
			fmt.Println("管理员账号创建成功！")
		},
	}

	return serverCmd
}
