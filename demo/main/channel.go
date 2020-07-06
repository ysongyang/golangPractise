package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("值=%v，原地址=%v\n", intChan, &intChan)
	intChan <- 10
	num := 21
	intChan <- num
	intChan <- 15
	fmt.Printf("inchan长度=%v,容量=%v\n", len(intChan), cap(intChan))
	n2 := <-intChan
	fmt.Printf("n2的值=%v\n", n2)
	n3 := <-intChan
	fmt.Printf("n3的值=%v\n", n3)
	n4 := <-intChan
	fmt.Printf("n4的值=%v\n", n4)
}
