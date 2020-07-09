package main

import (
	"fmt"
)

var (
	userId  int    //用户id
	userPwd string //用户密码
)

func loginMain() {
	var num int
	var loop = true //菜单是否显示
	for loop {
		fmt.Println("---------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t 1、登陆聊天室")
		fmt.Println("\t\t 2、注册用户")
		fmt.Println("\t\t 3、退出系统")
		fmt.Print("请选择（1—3）：")
		fmt.Scanf("%d\n", &num)
		switch num {
		case 1:
			showMenu(num)
			loop = false
		case 2:
			showMenu(num)
			loop = false
		case 3:
			fmt.Println("退出系统...")
			loop = false
			//os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入~")
		}
	}
	//根据用户的输入显示新的菜单
}

func showMenu(num int) {
	if num == 1 {
		fmt.Println("请输入用户的ID号：")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码：")
		fmt.Scanf("%s\n", &userPwd)
		if err := Login(userId, userPwd); err != nil {
			fmt.Println("login error", err)
		} else {
			fmt.Println("login success")
		}
	} else if num == 2 {
		fmt.Println("用户要注册")
	}
}

func main() {
	loginMain()
}
