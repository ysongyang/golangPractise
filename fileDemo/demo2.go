package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "e:/417法令模板.txt"

	if content, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println("readFile err = ", err)
	} else {
		fmt.Printf("%s", string(content))
	}
}
