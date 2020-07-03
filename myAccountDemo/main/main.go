package main

import (
	"Project/myAccountDemo/logic"
	"fmt"
)

func main() {

	fmt.Println("面向对象方法实现~~")
	logic.NewFamilyAccount().ShowMenu()
}
