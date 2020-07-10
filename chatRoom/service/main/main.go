package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"golangPractise/chatRoom/common/utils"
	"golangPractise/chatRoom/service/model"
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
		fmt.Println(err)
		return
	}
}

var RedisPool *redis.Pool

func initUserDao() {
	model.MyUserDao = model.NewUserDao(RedisPool)
}

func main() {
	RedisPool = utils.RedisPool() //初始化redis连接池
	initUserDao()
	fmt.Println("服务器已启动...正在监听8889端口")
	if listen, err := net.Listen("tcp", "0.0.0.0:8889"); err != nil {
		fmt.Println("listen error", err)
		return
	} else {
		defer listen.Close()
		for {
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

}
