package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().Unix()
	var flag bool
	for num := 1; num <= 80000; num++ {
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

		}
	}
	end := time.Now().Unix()
	fmt.Println("使用普通方法所耗费的时间(秒)：", end-start)
	fmt.Println("主线程完成...")
}
