package model

import (
	"fmt"
	"golangPractise/webApp/utils"
	"time"
)

//User 结构体
type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"-"` // "-"不会导出到JSON中
}

//添加User 方法一

func (user *User) AddUser() error {
	//sql语句
	sqlStr := "insert INTO users(username,password,email,mobile,created_at) values(?,?,?,?,?)"
	//预编译
	stmt, err := utils.Db.Prepare(sqlStr)
	//utils.Db.Close()
	if err != nil {
		fmt.Println("预编译出错：", err)
		return err
	}
	//执行
	_, errs := stmt.Exec(user.Username, user.Password, user.Email, user.Mobile, time.Now().Format("2006-01-02 15:04:05"))
	if errs != nil {
		fmt.Println("执行出错：", errs)
		return errs
	}
	return nil
}

//AddUser2 添加User 方法二

func (user *User) AddUser2() error {
	//sql语句
	sqlStr := "insert INTO users(username,password,email,mobile,created_at) values(?,?,?,?,?)"
	//预执行
	_, err := utils.Db.Exec(sqlStr, user.Username, user.Password, user.Email, user.Mobile, time.Now().Format("2006-01-02 15:04:05"))
	//utils.Db.Close()
	if err != nil {
		fmt.Println("执行出错：", err)
		return err
	}
	return nil
}

//根据id查询用户
func (user *User) GetUserById(userId int) (*User, error) {
	sql := "select id,username,password,email,mobile,created_at from users where id = ? "

	row := utils.Db.QueryRow(sql, userId)
	var id int
	var username string
	var password string
	var email string
	var mobile string
	var created_at string
	err := row.Scan(&id, &username, &password, &email, &mobile, &created_at)
	if err != nil {
		return nil, err
	}
	u := &User{
		Id:        id,
		Username:  username,
		Password:  password,
		Email:     email,
		Mobile:    mobile,
		CreatedAt: created_at,
	}
	return u, nil
}

func (user *User) GetUsers() ([]*User, error) {
	sql := "select id,username,password,email,mobile,created_at from users"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}

	//创建一个user切片接收数据
	var users []*User

	for rows.Next() {
		var id int
		var username string
		var password string
		var email string
		var mobile string
		var created_at string
		err := rows.Scan(&id, &username, &password, &email, &mobile, &created_at)
		if err != nil {
			return nil, err
		}
		users = append(users, &User{
			Id:        id,
			Username:  username,
			Password:  password,
			Email:     email,
			Mobile:    mobile,
			CreatedAt: created_at,
		})
	}
	return users, nil
}
