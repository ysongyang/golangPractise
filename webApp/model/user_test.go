package model

import (
	"fmt"
	"testing"
)

//测试之前可以执行的其他操作
func TestMain(m *testing.M) {
	fmt.Println("执行测试之前执行的...")
	m.Run()
}

//通过t.Run 执行子函数 进行测试

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法...")
	//t.Run("测试添加用户", testAddUser)
	//t.Run("测试查询用户", testGetUserById)
	t.Run("测试查询所有用户", testGetUsers)
}

func testAddUser(t *testing.T) {
	fmt.Println("测试AddUser方法...")
	user := &User{
		Username: "admin",
		Password: "123456",
		Email:    "admin@qq.com",
		Mobile:   "18675618521",
	}
	user.AddUser()

	user2 := &User{
		Username: "demo",
		Password: "123456",
		Email:    "admin@qq.com",
		Mobile:   "18675618521",
	}
	user2.AddUser2()
}

func testGetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录")
	user := &User{}
	user, err := user.GetUserById(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录")
	user := &User{}
	users, err := user.GetUsers()
	if err != nil {
		fmt.Println(err)
	}
	for index, val := range users {
		fmt.Printf("第%d个用户是%v \n", index+1, val)
	}
}
