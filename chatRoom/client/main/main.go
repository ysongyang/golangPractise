package main

import (
	"fmt"
	"golangPractise/chatRoom/client/process"
	"os"
)

var (
	userId  int    //用户id
	userPwd string //用户密码
)

func loginMain() {
	var num int
	for {
		fmt.Println("---------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t 1、登陆聊天室")
		fmt.Println("\t\t 2、注册用户")
		fmt.Println("\t\t 3、退出系统")
		fmt.Print("请选择（1—3）：")
		fmt.Scanf("%d\n", &num)
		switch num {
		case 1:
			fmt.Println("请输入用户的ID号：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码：")
			fmt.Scanf("%s\n", &userPwd)
			userLoginProcess := &process.UserLoginProcess{}
			err := userLoginProcess.Login(userId, userPwd)
			if err != nil {
				fmt.Println(err)
			}

		case 2:
		case 3:
			fmt.Println("退出系统...")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入~")
		}
	}
	//根据用户的输入显示新的菜单
}

func main() {
	loginMain()
}