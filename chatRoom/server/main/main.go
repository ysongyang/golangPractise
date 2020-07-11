package main

import (
	"fmt"
	"golangPractise/chatRoom/server/model"
	"net"
)

//处理和客户端通讯
func process(conn net.Conn) {
	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}
	err := processor.SynProcess()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误：", err)
		return
	}
}

func init() {
	RedisPool() //初始化redis连接池
	initUserDao()
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(redisPoolDb)
}

func main() {
	fmt.Println("服务器已启动...\n正在监听8899端口...")
	listen, err := net.Listen("tcp", "0.0.0.0:8899")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen error", err)
		return
	}
	for {
		fmt.Println("等待客户端来链接服务器.....")
		// 等待连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error", err)
			return
		}
		//链接成功则启动协程和客户端保持数据通讯
		go process(conn)
	}
}
