package process

import (
	"encoding/json"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
)

//客户端发送消息

type SmsProcess struct {
}

//群聊
func (sms *SmsProcess) SendGroupMes(content string) (err error) {
	//创建通讯消息体
	var msg message.Message
	msg.Type = message.SmsMsgType

	fmt.Println("当前用户ID：", curUser.UserId)

	//创建smsMes实例
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.User.UserId = curUser.UserId
	smsMsg.User.UserStatus = curUser.UserStatus

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("SendGroupMes json Marshal error", err.Error())
		return
	}
	msg.Data = string(data)

	//序列化msg
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("SendGroupMes json Marshal error", err.Error())
		return
	}
	//发送数据
	//创建一个Transfer实力
	transfer := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes WritePkg error", err.Error())
		return
	}
	return
}
