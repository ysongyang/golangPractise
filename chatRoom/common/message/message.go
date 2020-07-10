package message

import "golangPractise/chatRoom/service/model"

//服务器和客户端共用的消息体

const (
	LoginMsgType    = "LoginMsg" //登录类型
	LoginResMsgType = "LoginResMsg"
	RegMsgType      = "RegMsg" //
	ResponseMsgType = "ResponseMsg"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//定义登录消息体
type LoginMsg struct {
	UserId  int    `json:"user_id"`
	UserPwd string `json:"user_pwd"`
}

//登录返回消息体
type LoginResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//注册消息体
type RegMsg struct {
	user model.User `json:"user"`
}

//消息返回的消息体
type ResponseMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  string `json:"data"`
}
