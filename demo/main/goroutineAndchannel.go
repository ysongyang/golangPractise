package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		fmt.Println("writeData i =", i)
		intChan <- i
	}
	//关闭管道后，方便readData里对管道进行循环读
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		if v, ok := <-intChan; !ok {
			break
		} else {
			time.Sleep(time.Second)
			fmt.Println("readData v = ", v)
		}

	}

	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	//写协程
	go writeData(intChan)
	//读协程
	go readData(intChan, exitChan)

	for {
		if _, ok := <-exitChan; !ok {
			break
		}
	}
}
