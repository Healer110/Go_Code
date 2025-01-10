package main

import (
	"chatRoom_system/client/process"
	"fmt"
	"os"
)

// 定义两个变量，表示用户名密码,以及用户昵称
var (
	userId   int
	userPwd  string
	username string
)

func firstMenu() {
	fmt.Println("------------------- 欢迎登录多人聊天系统 -------------------")
	fmt.Println("                     1 登录系统")
	fmt.Println("                     2 注册用户")
	fmt.Println("                     3 退出系统")
	fmt.Println("请选择(1-3):")
}

func main() {
	// 接收用户的输入
	var key int
	// 判断是否还继续显示菜单
	for {
		firstMenu()
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			// 完成登录，创建UserProcess实例
			up := process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("error:", err)
			}

		case 2:
			fmt.Println("请输入用户的id:")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入用户的名字:")
			fmt.Scanln(&username)
			// 完成注册请求
			up := process.UserProcess{}
			err := up.Register(userId, userPwd, username)
			if err != nil {
				fmt.Println("[client/main/main.go] Register user error: ", err)
			}

		case 3:
			fmt.Println("退出系统...")
			os.Exit(0)
			// loop = false
		default:
			fmt.Println("输入错误，请重新输入")
		}

	}

}
