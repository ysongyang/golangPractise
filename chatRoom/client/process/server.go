package process

import (
	"encoding/json"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	"golangPractise/chatRoom/server/model"
	"net"
	"os"
)

//显示登陆菜单、保持服务端通讯

//显示登陆成功后的界面
func ShowMenu(data string) {
	var user model.User
	var num int
	var content string
	sProcess := &SmsProcess{}
	json.Unmarshal([]byte(data), &user)
	fmt.Printf("------------------恭喜< %s >登陆成功------------------\n", user.UserName)
	fmt.Println("\t\t 1、显示在线用户列表")
	fmt.Println("\t\t 2、发送信息")
	fmt.Println("\t\t 3、信息列表")
	fmt.Println("\t\t 4、退出系统")
	fmt.Println("请选择（1—4）：")
	fmt.Scanf("%d\n", &num)
	switch num {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你要发送的内容：")
		fmt.Scanf("%s\n", &content)
		sProcess.SendGroupMes(content)
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
		//fmt.Println("message = ", res)

		switch res.Type {
		case message.NotifyUserStatusMsgType:
			//用户状态发生变化了
			//1.取出NotifyUserStatusMsg
			var notifyMsg message.NotifyUserStatusMsg
			json.Unmarshal([]byte(res.Data), &notifyMsg)
			updateUserStatus(&notifyMsg)
		//2.把用户的信息，状态保存到客户端的map[int]User
		case message.SmsMsgType: //处理服务端发送过来的群发消息
			outputGroupMsg(&res)
		default:
			fmt.Println("服务器返回了一个未知的消息类型", res.Type)
		}
	}
}
