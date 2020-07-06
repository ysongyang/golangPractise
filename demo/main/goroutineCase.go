package main

import (
	"fmt"
)

func main() {

	intChan := make(chan int, 10)
	strChan := make(chan string, 5)

	for i := 1; i <= 10; i++ {
		intChan <- i
	}

	for i := 1; i <= 5; i++ {
		//strChan <- strconv.Itoa(i) + "__hello world"
		strChan <- "hello " + fmt.Sprintf("%d", i)
	}
	for {
		select {
		case v := <-intChan:
			fmt.Println("从inChan读取数据 = ", v)
		case s := <-strChan:
			fmt.Println("从strChan读取数据 = ", s)
		default:
			fmt.Println("管道读取完成...")
			return
		}
	}
}
