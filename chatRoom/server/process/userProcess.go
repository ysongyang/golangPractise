package processes

import (
	"encoding/json"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	"golangPractise/chatRoom/server/model"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int //增加一个字段，标识conn属于哪个用户
}

//服务端通知所有在线的用户
func (userProcess *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onLineUser 一个一个的发送NotifyUserStatusMsg
	for uid, uProcess := range userMgr.onLineUsers {
		//如果是客户端自己 跳出
		if uid == userId {
			continue
		}
		//通知
		uProcess.NotifyMeOnline(userId)
	}
}

//给大家通知客户端上线
func (userProcess *UserProcess) NotifyMeOnline(userId int) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType

	var notifyMsg message.NotifyUserStatusMsg

	notifyMsg.UserId = userId
	notifyMsg.Status = message.UserOnline

	//通知消息体序列化
	data, err := json.Marshal(notifyMsg)
	if err != nil {
		fmt.Println("NotifyMeOnline json Marshal error ", err)
		return
	}
	msg.Data = string(data)

	//客户端和服务端通讯的消息体 序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("NotifyMeOnline msg json Marshal error ", err)
		return
	}
	//fmt.Println("msg =", msg)
	// msg = {NotifyUserStatusMsg {"user_id":100,"status":1}}
	//发送数据
	//创建一个Transfer实力
	transfer := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	//fmt.Println("data =", string(data))
	//data = {"type":"NotifyUserStatusMsg","data":"{\"user_id\":1,\"status\":1}"}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline WritePkg error ", err)
		return
	}
}

//注册的逻辑
func (userProcess *UserProcess) ServiceProcessRegister(mes *message.Message) (err error) {
	var regMsg message.RegMsg
	//先从mes中取出mes.Data 并直接反序列化成RegMsg
	err = json.Unmarshal([]byte(mes.Data), &regMsg)
	if err != nil {
		fmt.Println("ServiceProcessRegister json Unmarshal error", err)
		return
	}

	var msg message.Message

	msg.Type = message.ResponseMsgType

	var responseMsg message.ResponseMsg

	//这里需要传结构体
	err = model.MyUserDao.Register(&regMsg.User)
	if err != nil {
		responseMsg.Error = err.Error()
		responseMsg.Code = 201

	} else {
		responseMsg.Code = 200
		responseMsg.Data = mes.Data
		responseMsg.Error = fmt.Sprintf("%s", "register success")
	}

	data, err := json.Marshal(responseMsg)
	if err != nil {
		fmt.Println("ServiceProcessRegister responseMsg json Marshal error", err)
		return
	}

	msg.Data = string(data) //切片转字符串

	//对 resMsg序列化 发送给 客户端
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("ServiceProcessRegister msg json Marshal error", err)
		return
	}
	//发送数据
	//创建一个Transfer实力
	transfer := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = transfer.WritePkg(data)
	return
}

//登录的逻辑
func (userProcess *UserProcess) ServiceProcessLogin(mes *message.Message) (err error) {
	//先从mes中取出mes.Data 并直接反序列化成LoginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(mes.Data), &loginMsg)
	if err != nil {
		fmt.Println("serviceProcessLogin json.Unmarshal error", err)
		return
	}

	var msg message.Message
	msg.Type = message.LoginResMsgType

	var loginResMsg message.LoginResMsg

	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		loginResMsg.Error = err.Error()
		loginResMsg.Code = 201

	} else {
		userJson, _ := json.Marshal(user)
		loginResMsg.Code = 200
		loginResMsg.Data = string(userJson)
		loginResMsg.Error = "login success"
		//将登录成的用户id赋给userProcess.UserId
		userProcess.UserId = loginMsg.UserId
		//登录成功，把该用户放入到userMgr
		userMgr.AddOnlineUser(userProcess)

		//通知其他在线的用户，客户端上线了
		userProcess.NotifyOthersOnlineUser(loginMsg.UserId)

		//将当前在线用户的id放入 responseMsg.UsersId 对 userMgr.onLineUsers 进行遍历
		for id, _ := range userMgr.onLineUsers {
			loginResMsg.UsersId = append(loginResMsg.UsersId, id)
		}

		fmt.Println(user, "登录成功")
	}

	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("serviceProcessLogin loginResMsg json Marshal error ", err)
		return
	}
	msg.Data = string(data) //切片转字符串

	//对 resMsg序列化 发送给 客户端
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("serviceProcessLogin msg json Marshal error ", err)
		return
	}
	//发送数据
	//创建一个Transfer实力
	transfer := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = transfer.WritePkg(data)
	return
}
