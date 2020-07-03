package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("e:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	/*if err = file.Close(); err != nil {
		fmt.Println("close file err = ", err)
	}*/
	defer file.Close()
	/*const (

		defaultBufSize = 4096 //默认缓冲区
	)*/

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') //读到换行符结束一次
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
}
