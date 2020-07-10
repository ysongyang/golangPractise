package process

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	"net"
)

type UserLoginProcess struct {
}

//登录的校验
func (userLogin *UserLoginProcess) Login(userId int, userPwd string) (err error) {
	//链接服务器端
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Login net Dial error", err)
		return errors.New(errText)
	}
	defer conn.Close() //延时关闭
	//发送消息给服务器
	var msg message.Message
	msg.Type = message.LoginMsgType

	//创建LoginMsg结构体
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	//序列化loginMsg
	data, err := json.Marshal(loginMsg)
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Login loginMsg json Marshal error", err)
		return errors.New(errText)
	}
	msg.Data = string(data)

	//序列化msg
	dataMsg, error := json.Marshal(msg)
	if error != nil {
		errText := fmt.Sprintf("%s:%s", "Login msg json Marshal error", err)
		return errors.New(errText)
	}

	//发送消息长度
	//1. 先获取dataMsg长度 转成一个表示长度的切片
	var buf [4]byte
	//长度转成切片
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(dataMsg)))
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		errText := fmt.Sprintf("%s:%s", "Login conn.Write buf error", err)
		return errors.New(errText)
	}
	fmt.Printf("客户端发送数据消息长度=%d 内容=%s\n", len(dataMsg), string(dataMsg))

	//发送消息体
	_, errs := conn.Write(dataMsg)
	if errs != nil {
		errText := fmt.Sprintf("%s:%s", "Login conn.Write dataMsg error", err)
		return errors.New(errText)
	}

	/*time.Sleep(time.Second * 10)
	fmt.Println("客户端休眠了10秒关闭")*/

	//创建一个Transfer实力
	transfer := &utils.Transfer{
		Conn: conn,
	}
	//接收服务器返回的数据
	res, err := transfer.ReadPkg()
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Login ReadPkg error", err)
		return errors.New(errText)
	}
	fmt.Printf("收到服务端数据 %v\n", res)
	var loginResMsg message.LoginResMsg
	//将res的Data反序列化成loginResMsg
	err = json.Unmarshal([]byte(res.Data), &loginResMsg)
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Login loginResMsg json Unmarshal error", err)
		return errors.New(errText)
	}
	if loginResMsg.Code == 200 {
		//启动一个隐藏的协程，该协程保持和服务端的通讯，如果服务端有数据流则进行推送给客户端
		go ProcessesServiceMsg(conn)
		//显示登陆成功的菜单
		for {
			ShowMenu()
		}
	}
	fmt.Println(loginResMsg.Error)
	return
}
