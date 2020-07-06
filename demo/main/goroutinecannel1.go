package main

import (
	"fmt"
	"time"
)

func putData(intChan chan int, n int) {

	for i := 1; i <= n; i++ {
		intChan <- i
	}

	close(intChan)
}

func primeData(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		//如果管道取不到数据
		if !ok {
			break
		}
		flag = true
		//判断num 是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	//
	exitChan <- true
	fmt.Println("协程完成工作... ...")
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()

	go putData(intChan, 80000)

	go primeData(intChan, primeChan, exitChan)
	go primeData(intChan, primeChan, exitChan)
	go primeData(intChan, primeChan, exitChan)
	go primeData(intChan, primeChan, exitChan)

	//这里处理主线程
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程所耗费的时间(秒)：", end-start)
		close(primeChan)
	}()
	//打印素数数据
	for {
		//取数据
		if _, ok := <-primeChan; !ok {
			break
		} else {
			//fmt.Println("素数 = ", v)
		}
	}
	fmt.Println("主线程完成...")
}
