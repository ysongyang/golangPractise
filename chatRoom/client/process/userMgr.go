package process

import (
	"fmt"
	clientModel "golangPractise/chatRoom/client/model"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/server/model"
)

//客户端的User
var onlineUsers map[int]*model.User = make(map[int]*model.User, 1024)

var curUser clientModel.CurUser //在用户登录成功完成初始化

func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, user := range onlineUsers {
		fmt.Printf("用户ID：%d\t用户名：%s\t\n", id, user.UserName)
	}
}

//更新用户状态
//处理返回的NotifyUserStatusMsg
func updateUserStatus(msg *message.NotifyUserStatusMsg) {
	user, ok := onlineUsers[msg.UserId]
	//如果不存在
	if !ok {
		//创建一个User实例
		user = &model.User{
			UserId:     msg.UserId,
			UserStatus: msg.Status,
		}
	}
	//更改状态
	user.UserStatus = msg.Status
	onlineUsers[msg.UserId] = user
	outputOnlineUser()
}
