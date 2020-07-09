package main

import "fmt"

//登录的校验
func Login(userId int, userPwd string) (err error) {
	fmt.Printf("你输入的用户ID %d，密码：%s\n", userId, userPwd)
	return nil
}
