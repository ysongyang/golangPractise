package main

import (
	"fmt"
	"golangPractise/customerPro/service"
)

type customerView struct {
	//定义字段
	key             string //控制台输入的编号
	loop            bool   //表示是否循环显示菜单
	customerService *service.CustomerService
}

//主菜单方法
func (c *customerView) showMenu() {
	for {
		fmt.Println("---------------------客户信息管理软件---------------------")
		fmt.Println("                     1、添加客户                     ")
		fmt.Println("                     2、修改客户                     ")
		fmt.Println("                     3、删除客户                     ")
		fmt.Println("                     4、客户列表                     ")
		fmt.Println("                     5、退   出                     ")
		fmt.Print("请选择（1-5）：")
		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			c.list()
		case "5":
			c.loop = true

		default:
			fmt.Println("您的输入有误，请重新输入！")
		}

		if c.loop {
			fmt.Println("您退出了客户关系管理系统...")
			break
		}

	}
}

//客户列表
func (c *customerView) list() {
	fmt.Println("-------------------------------客户列表-------------------------------")
	//获取当前所有的客户信息
	listData := c.customerService.List()
	if len(listData) > 0 {
		fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
		for i := 0; i < len(listData); i++ {
			//调用model中的GetInfo()方法进行数据显示
			fmt.Println(listData[i].GetInfo())
		}
		println()
		println()
	} else {
		fmt.Println("暂无数据...")
	}
}

func main() {

	cv := customerView{
		key:             "",
		loop:            false,
		customerService: service.NewCustomerService(),
	}
	cv.showMenu()
}
