package processes

import (
	"encoding/json"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	"net"
)

//群发消息

type SmsProcess struct {
}

//转发消息

func (sp *SmsProcess) SendGroupMsg(msg *message.Message) {

	//取出message的内容
	var smsMsg message.SmsMsg

	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("SendGroupMsg json Unmarshal Error", err.Error())
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("SendGroupMsg json Marshal Error", err.Error())
		return
	}

	//遍历服务器端的map  onLineUsers map[int]*UserProcess
	for id, uProcess := range userMgr.onLineUsers {
		if id == smsMsg.UserId {
			continue
		}
		sp.SendMsg(data, uProcess.Conn)
	}
}

func (sp *SmsProcess) SendMsg(data []byte, conn net.Conn) {
	transfer := &utils.Transfer{
		Conn: conn, //取客户端的conn
	}
	//发送数据
	err := transfer.WritePkg(data)
	if err != nil {
		fmt.Println("发送消息失败", err.Error())
	}
}
