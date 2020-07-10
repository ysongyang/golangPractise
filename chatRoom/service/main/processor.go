package main

import (
	"errors"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	processes "golangPractise/chatRoom/service/process"
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
		if err == io.EOF {
			//fmt.Println("SynProcess exit...")
			//return
			errText := fmt.Sprintf("%s:%s", "SynProcess exit...", err)
			return errors.New(errText)
		} else if err != nil {
			//fmt.Println("SynProcess ReadPkg fail:", err)
			errText := fmt.Sprintf("%s:%s", "SynProcess ReadPkg error", err)
			return errors.New(errText)
		}
		err = processor.ServiceProcessMsg(&msg)
		if err != nil {
			//fmt.Println("SynProcess ServiceProcessMsg error:", err)
			errText := fmt.Sprintf("%s:%s", "SynProcess ServiceProcessMsg error", err)
			return errors.New(errText)
		}
		fmt.Println("SynProcess success:", msg)
		return nil
	}

}

//根据客户端发送的消息类型调用对应的函数
func (processor *Processor) ServiceProcessMsg(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMsgType:
		userProcess := &processes.UserProcess{
			Conn: processor.Conn,
		}
		err = userProcess.ServiceProcessLogin(mes)
		//处理登录的逻辑
	case message.RegMsgType:
	//注册的逻辑
	default:
		errText := fmt.Sprintf("%s:%s", "ServiceProcessMsg不存在的消息类型", err)
		return errors.New(errText)
	}
	return
}
