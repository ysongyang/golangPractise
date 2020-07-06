package main

import (
	"fmt"
	"golangPractise/utils"
)

func main() {
	str := []byte("12fff我是ww.topgoer.com的站长枯藤")
	pwd, _ := utils.EncryptionCode(str)
	fmt.Println("加密后：", pwd)
	bytes, _ := utils.DecryptCode(pwd)
	fmt.Println("解密后：", string(bytes))

}
