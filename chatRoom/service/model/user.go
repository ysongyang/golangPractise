package model

//用户结构体

type User struct {
	//注意这里定义的json tag 要跟 message.go 里的 json tag 对应
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
}
