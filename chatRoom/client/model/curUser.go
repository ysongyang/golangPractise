package model

import (
	"golangPractise/chatRoom/server/model"
	"net"
)

//全局的结构体变量
type CurUser struct {
	Conn net.Conn
	model.User
}
