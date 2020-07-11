package process

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	"net"
	"time"
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
	dataMsg, err := json.Marshal(msg)
	if err != nil {
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
	var responseMsg message.ResponseMsg
	//将res的Data反序列化成loginResMsg
	err = json.Unmarshal([]byte(res.Data), &responseMsg)
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Login loginResMsg json Unmarshal error", err)
		return errors.New(errText)
	}
	if responseMsg.Code == 200 {
		//启动一个隐藏的协程，该协程保持和服务端的通讯，如果服务端有数据流则进行推送给客户端
		go ProcessesServiceMsg(conn)
		//显示登陆成功的菜单
		for {
			ShowMenu(res.Data)
		}
	}
	fmt.Println(responseMsg.Error)
	return
}

//注册
func (userLogin *UserLoginProcess) Register(userId int, userPwd string, userName string, phone string) (err error) {
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Register net Dial error", err)
		return errors.New(errText)
	}
	defer conn.Close() //延时关闭
	//发送消息给服务器
	var msg message.Message
	msg.Type = message.RegMsgType

	//创建RegMsg结构体
	var regMsg message.RegMsg
	regMsg.User.UserId = userId
	regMsg.User.UserPwd = userPwd
	regMsg.User.UserName = userName
	regMsg.User.Phone = phone
	regMsg.User.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	//序列化RegMsg
	data, err := json.Marshal(regMsg)
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Register loginMsg json Marshal error", err)
		return errors.New(errText)
	}
	//data切片转换成string
	msg.Data = string(data)

	//序列化Message
	dataMsg, err := json.Marshal(msg)
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Register msg json Marshal error", err)
		return errors.New(errText)
	}

	//创建一个Transfer实力
	transfer := &utils.Transfer{
		Conn: conn,
	}
	//发送数据到服务器
	err = transfer.WritePkg(dataMsg)
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Register WritePkg error", err)
		return errors.New(errText)
	}
	//fmt.Printf("服务端数据发送成功 %v\n", dataMsg)

	//接收服务器返回的数据
	res, err := transfer.ReadPkg()
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Login ReadPkg error", err)
		return errors.New(errText)
	}
	fmt.Printf("收到服务端数据 %v\n", res)

	var resMsg message.ResponseMsg
	//将res的Data反序列化成loginResMsg
	err = json.Unmarshal([]byte(res.Data), &resMsg)
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "Register ResponseMsg json Unmarshal error", err)
		return errors.New(errText)
	}
	if resMsg.Code == 200 {
		//注册成功
		fmt.Println("注册成功")
		//os.Exit(0)
	} else {
		fmt.Println(resMsg.Error)
		//os.Exit(0)
	}

	return
}
