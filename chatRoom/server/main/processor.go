package main

import (
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	processes "golangPractise/chatRoom/server/process"
	"io"
	"net"
)

type Processor struct {
	Conn      net.Conn
	ProcessId int64
}

//处理和客户端通讯
func (processor *Processor) SynProcess() (err error) {
	//循环读取客户端发送的消息
	for {
		//创建一个Transfer实例
		transfer := &utils.Transfer{
			Conn: processor.Conn,
		}
		msg, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出..")
				return err
			} else {
				fmt.Println("SynProcess ReadPkg error", err)
				return err
			}
		}

		err = processor.ServiceProcessMsg(&msg)
		if err != nil {
			fmt.Println("SynProcess ServiceProcessMsg error", err)
			return err
		}
	}
}

//根据客户端发送的消息类型调用对应的函数
func (processor *Processor) ServiceProcessMsg(mes *message.Message) (err error) {
	//看看是否能接收到客户端发送的群发的消息
	//fmt.Println("ServiceProcessMsg msg ", mes)
	switch mes.Type {
	case message.LoginMsgType: //处理登录的逻辑
		userProcess := &processes.UserProcess{
			Conn: processor.Conn,
		}
		userProcess.ServiceProcessLogin(mes)

	case message.RegMsgType: //注册的逻辑
		userProcess := &processes.UserProcess{
			Conn: processor.Conn,
		}
		userProcess.ServiceProcessRegister(mes)

	case message.SmsMsgType: //发送消息的逻辑
		smsProcess := processes.SmsProcess{}
		smsProcess.SendGroupMsg(mes)
	default:
		fmt.Println("ServiceProcessMsg不存在的消息类型", err)
		return
	}
	return
}
