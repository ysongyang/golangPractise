package main

import (
	"fmt"
	"log"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//创建切片
		buf := make([]byte, 1024)
		//fmt.Println("服务端等待客户端的数据,客户端IP：", conn.RemoteAddr())
		n, err := conn.Read(buf) //从conn 读取数据 ，等待客户端发送数据，如果客户端没有Write，这里将会阻塞
		if err != nil {
			fmt.Println("客户端退出...")
			break
		}
		fmt.Printf("收到客户端 %v 发来的内容 %v", conn.RemoteAddr(), string(buf[:n]))
	}

}

func main() {

	fmt.Println("服务器开始监听...")
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("listen = %v\n", l)
	defer l.Close()
	for {
		// 等待连接
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go process(conn)
		/*go func(conn net.Conn) {
			// Echo all incoming data.
			// io.Copy(conn, conn)
			// Shut down the connection.
			conn.Close()
		}(conn)*/
	}
}
