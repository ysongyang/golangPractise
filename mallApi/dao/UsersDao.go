package dao

import (
	"golangPractise/mallApi/model"
	"golangPractise/mallApi/utils"
	"log"
	"time"
)

func InsertUser(user *model.Users) (*model.Users, error) {
	//sql语句
	sqlStr := "insert INTO users(username,password,email,mobile,created_at) values(?,?,?,?,?)"
	//预执行
	res, err := utils.Db.Exec(sqlStr, user.Username, user.Password, user.Email, user.Mobile, time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		log.Println("insert user fail：", err)
		return nil, err
	}
	id, _ := res.LastInsertId()
	user, _ = GetUserById(int(id))
	return user, nil
}

//根据id查询用户
func GetUserById(userId int) (*model.Users, error) {
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
	u := &model.Users{
		Id:        id,
		Username:  username,
		Password:  password,
		Email:     email,
		Mobile:    mobile,
		CreatedAt: created_at,
	}
	return u, nil
}
