package model

import "errors"

//自定义错误信息

var (
	ERROR_USER_NOTEXIST = errors.New("用户不存在")
	ERROR_USER_EXISTS   = errors.New("用户已经存在")
	ERROR_USER_PASSWORD = errors.New("密码不正确")
)
