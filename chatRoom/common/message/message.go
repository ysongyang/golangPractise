package message

import "golangPractise/chatRoom/server/model"

//服务器和客户端共用的消息体

const (
	LoginMsgType            = "LoginMsg" //登录类型
	LoginResMsgType         = "LoginResMsg"
	RegMsgType              = "RegMsg" //
	ResponseMsgType         = "ResponseMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMsgType              = "SmsMsg"
)

//常量(用户状态)
const (
	UserOffline    = iota //离线		0
	UserOnline            //上线		1
	UserBusyStatus        //繁忙		2
)

//客户端和服务端通讯的消息体
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//登录消息体
type LoginMsg struct {
	UserId  int    `json:"user_id"`
	UserPwd string `json:"user_pwd"`
}

//注册消息体
type RegMsg struct {
	User model.User `json:"user"`
}

//消息消息体
type SmsMsg struct {
	Content string `json:"content"`
	model.User
}

//配合服务端通知用户上线状态消息体(服务器推送给客户端)
type NotifyUserStatusMsg struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

//登录返回的消息体
type LoginResMsg struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Error   string `json:"error"`
	UsersId []int  //登录成功后返回所有在线的userId
}

//消息返回的消息体（通用）
type ResponseMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  string `json:"data"`
}
