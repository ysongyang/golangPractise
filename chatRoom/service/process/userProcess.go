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

	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType

	var loginResMsg message.LoginResMsg

	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		loginResMsg.Error = err.Error()
		loginResMsg.Code = 201

	} else {
		loginResMsg.Code = 200
		loginResMsg.Error = fmt.Sprintf("%s %s", user.UserName, "login success")
	}

	data, err := json.Marshal(loginResMsg)
	if err != nil {
		//fmt.Println("serviceProcessLogin loginResMsg json Marshal error", err)
		//return
		errText := fmt.Sprintf("%s:%s", "serviceProcessLogin loginResMsg json Marshal error", err)
		return errors.New(errText)
	}
	resMsg.Data = string(data) //切片转字符串

	//对 resMsg序列化 发送给 客户端
	data, err = json.Marshal(resMsg)
	if err != nil {
		//fmt.Println("serviceProcessLogin resMsg json Marshal error", err)
		//return
		errText := fmt.Sprintf("%s:%s", "serviceProcessLogin resMsg json Marshal error", err)
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
