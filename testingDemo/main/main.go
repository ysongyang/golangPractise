package main

import "fmt"

func addUpper(n int) int {
	res := 0

	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	if res := addUpper(10); res != 55 {
		fmt.Println("error!")
	} else {
		fmt.Println("ok")
	}

}
