package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	"golangPractise/chatRoom/server/model"
	"net"
	"time"
)

type UserProcessClient struct {
}

//登录的校验
func (userProcess *UserProcessClient) Login(userId int, userPwd string) (err error) {
	//链接服务器端
	conn, err := net.Dial("tcp", "0.0.0.0:8899")
	if err != nil {
		fmt.Println("Login net Dial error", err)
		return
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
		fmt.Println("json Marshal error", err)
		return
	}
	msg.Data = string(data)

	//序列化msg
	dataMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json Marshal error", err)
		return
	}

	//发送消息长度
	//1. 先获取dataMsg长度 转成一个表示长度的切片
	var buf [4]byte
	//长度转成切片
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(dataMsg)))
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("Login conn.Write buf error", err)
		return
	}
	//fmt.Printf("客户端发送数据消息长度=%d 内容=%s\n", len(dataMsg), string(dataMsg))

	//发送消息体
	_, errs := conn.Write(dataMsg)
	if errs != nil {
		fmt.Println("Login conn.Write dataMsg error", err)
		return
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
		fmt.Println("Login ReadPkg error", err)
		return
	}
	fmt.Println("收到服务端数据:", res)
	var loginResMsg message.LoginResMsg
	//将res的Data反序列化成loginResMsg
	err = json.Unmarshal([]byte(res.Data), &loginResMsg)
	if loginResMsg.Code == 200 {

		//初始化 curUser
		curUser.Conn = conn
		curUser.User.UserId = userId
		curUser.User.UserStatus = message.UserOnline

		//登录成功 显示在线列表
		fmt.Println("当前在线用户列表如下：")
		for _, uid := range loginResMsg.UsersId {
			//过滤到自己的uid
			if userId == uid {
				continue
			}
			fmt.Println("用户ID:\t", uid)
			user := &model.User{
				UserId:     uid,
				UserStatus: message.UserOnline,
			}
			onlineUsers[uid] = user
		}
		fmt.Printf("\n\n")
		//启动一个隐藏的协程，该协程保持和服务端的通讯，如果服务端有数据流则进行推送给客户端
		go ProcessesServiceMsg(conn)
		//显示登陆成功的菜单
		for {
			ShowMenu(loginResMsg.Data)
		}
	} else {
		fmt.Println(loginResMsg.Error)
	}
	return
}

//注册
func (userProcess *UserProcessClient) Register(userId int, userPwd string, userName string, phone string) (err error) {
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("Register net Dial error", err)
		return
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
	regMsg.User.UserStatus = message.UserOffline
	regMsg.User.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	//序列化RegMsg
	data, err := json.Marshal(regMsg)
	if err != nil {
		fmt.Println("Register loginMsg json Marshal error", err)
		return
	}
	//data切片转换成string
	msg.Data = string(data)

	//序列化Message
	dataMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Register msg json Marshal error", err)
		return
	}

	//创建一个Transfer实力
	transfer := &utils.Transfer{
		Conn: conn,
	}
	//发送数据到服务器
	err = transfer.WritePkg(dataMsg)
	if err != nil {
		fmt.Println("Register WritePkg error", err)
		return
	}
	//fmt.Printf("服务端数据发送成功 %v\n", dataMsg)

	//接收服务器返回的数据
	res, err := transfer.ReadPkg()
	if err != nil {
		fmt.Println("Login ReadPkg error", err)
		return
	}
	fmt.Printf("收到服务端数据 %v\n", res)

	var resMsg message.ResponseMsg
	//将res的Data反序列化成loginResMsg
	err = json.Unmarshal([]byte(res.Data), &resMsg)
	if err != nil {
		fmt.Println("Register ResponseMsg json Unmarshal error", err)
		return
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
