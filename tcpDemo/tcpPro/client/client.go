package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.30:8888")
	if err != nil {
		// handle error
		fmt.Println("handle error", err)
		return
	}
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	for {
		fmt.Print("请输入内容：")
		reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("err", err)
			return
		}
		reader = strings.Trim(reader, " \r\n")
		if reader == "exit" {
			fmt.Println("客户端退出了!")
			break
		}
		n, err := conn.Write([]byte(reader + "\n"))
		if err != nil {
			fmt.Println("conn Write err", err)
			return
		}
		fmt.Printf("客户端发送了%v 字节的数据\n", n)
	}

}
