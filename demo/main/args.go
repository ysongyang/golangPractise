package main

import (
	"flag"
	"fmt"
)

func main() {

	/*args := os.Args
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}*/

	//定义变量
	var user string
	var pwd string
	var host string
	var port int

	//接收命令行中 -u root 中的 root
	flag.StringVar(&user, "u", "root", "用户名,默认为root")
	flag.StringVar(&pwd, "p", "root", "密码,默认为root")
	flag.StringVar(&host, "h", "localhost", "地址,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认3306")

	//转换
	//来解析命令行参数写入注册的flag里。
	//解析之后，flag的值可以直接使用。如果你使用的是flag自身，它们是指针；如果你绑定到了某个变量，它们是值。
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v\n", user, pwd, host, port)
}
