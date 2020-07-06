package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func goroutineTest() {
	for i := 1; i <= 10; i++ {
		fmt.Println("goroutineTest hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	look.Lock() //加锁
	myMap[n] = res
	look.Unlock() //解锁
}

var (
	myMap = make(map[int]int, 10)
	look  sync.Mutex //全局变量互斥锁
)

func main() {
	/*fmt.Println("main")
	//go goroutineTest()

	cpuNum := runtime.NumCPU()
	fmt.Println("cpu的数量：", cpuNum)

	runtime.GOMAXPROCS(2)

	fmt.Println("cpu的数量：", cpuNum)

	for i := 1; i <= 10; i++ {
		fmt.Println("main hello golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}*/

	for i := 1; i < 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	look.Lock()
	for i := 0; i <= len(myMap); i++ {
		fmt.Printf("map[%v] = %v\n", i+1, myMap[i])
	}
	look.Unlock()

}
