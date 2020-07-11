package processes

import (
	"encoding/json"
	"errors"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"golangPractise/chatRoom/common/utils"
	"golangPractise/chatRoom/service/model"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//注册的逻辑
func (userProcess *UserProcess) ServiceProcessRegister(mes *message.Message) (err error) {
	var regMsg message.RegMsg
	//先从mes中取出mes.Data 并直接反序列化成RegMsg
	err = json.Unmarshal([]byte(mes.Data), &regMsg)
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "ServiceProcessRegister json Unmarshal error", err)
		return errors.New(errText)
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
		//fmt.Println("serviceProcessLogin loginResMsg json Marshal error", err)
		//return
		errText := fmt.Sprintf("%s:%s", "ServiceProcessRegister responseMsg json Marshal error", err)
		return errors.New(errText)
	}

	msg.Data = string(data) //切片转字符串

	//对 resMsg序列化 发送给 客户端
	data, err = json.Marshal(msg)
	if err != nil {
		//fmt.Println("serviceProcessLogin resMsg json Marshal error", err)
		//return
		errText := fmt.Sprintf("%s:%s", "ServiceProcessRegister msg json Marshal error", err)
		return errors.New(errText)
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
		//fmt.Println("serviceProcessLogin json.Unmarshal fail", err)
		errText := fmt.Sprintf("%s:%s", "serviceProcessLogin json.Unmarshal error", err)
		return errors.New(errText)
	}

	var msg message.Message
	msg.Type = message.LoginResMsgType

	var responseMsg message.ResponseMsg

	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		responseMsg.Error = err.Error()
		responseMsg.Code = 201

	} else {
		userJson, _ := json.Marshal(user)
		responseMsg.Code = 200
		responseMsg.Data = string(userJson)
		responseMsg.Error = fmt.Sprintf("%s", "login success")
	}

	data, err := json.Marshal(responseMsg)
	if err != nil {
		//fmt.Println("serviceProcessLogin loginResMsg json Marshal error", err)
		//return
		errText := fmt.Sprintf("%s:%s", "serviceProcessLogin loginResMsg json Marshal error", err)
		return errors.New(errText)
	}
	msg.Data = string(data) //切片转字符串

	//对 resMsg序列化 发送给 客户端
	data, err = json.Marshal(msg)
	if err != nil {
		//fmt.Println("serviceProcessLogin resMsg json Marshal error", err)
		//return
		errText := fmt.Sprintf("%s:%s", "serviceProcessLogin msg json Marshal error", err)
		return errors.New(errText)
	}
	//发送数据
	//创建一个Transfer实力
	transfer := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = transfer.WritePkg(data)
	return
}
