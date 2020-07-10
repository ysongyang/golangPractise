package process

import (
	"fmt"
	"golangPractise/chatRoom/common/utils"
	"net"
	"os"
)

//显示登陆菜单、保持服务端通讯

//显示登陆成功后的界面
func ShowMenu() {

	fmt.Println("---------------恭喜xxx登陆成功---------------")
	fmt.Println("\t\t 1、显示在线用户列表")
	fmt.Println("\t\t 2、发送信息")
	fmt.Println("\t\t 3、信息列表")
	fmt.Println("\t\t 4、退出系统")
	fmt.Print("请选择（1—4）：")
	var num int
	fmt.Scanf("%d\n", &num)
	switch num {
	case 1:
		fmt.Println("显示用户列表")
	case 2:
		fmt.Println("发送信息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("您的输入有误，请重新输入~")
	}
}

//和服务端保持通讯
func ProcessesServiceMsg(conn net.Conn) {
	//创建一个Transfer 实例
	transfer := &utils.Transfer{
		Conn: conn,
	}
	for {
		//fmt.Println("\n客户端正在等待服务器发送的消息")
		res, err := transfer.ReadPkg()
		if err != nil {
			//fmt.Println("ProcessesServiceMsg ReadPkg error", err)
			return
		}
		//如果读取到数据
		fmt.Println("message = ", res)
	}
}
