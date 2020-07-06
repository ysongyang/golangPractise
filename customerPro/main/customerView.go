package main

import (
	"fmt"
	"golangPractise/customerPro/model"
	"golangPractise/customerPro/service"
	"golangPractise/utils"
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
			c.add()
		case "2":
			c.edit()
		case "3":
			c.delete()
		case "4":
			c.list()
		case "5":
			c.exit()
		default:
			fmt.Println("您的输入有误，请重新输入！")
		}

		if c.loop {
			fmt.Println("成功退出客户关系管理系统")
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

//校验手机号
func verifyPhone(phone string) string {
	mobile := ""
	flag := utils.VerifyMobileFormat(phone)
	if flag == false {
		fmt.Println("请输入有效的手机号！")
		fmt.Print("手机号：")
		fmt.Scanln(&mobile)
		verifyPhone(mobile)
	} else {
		mobile = phone
	}
	return mobile
}

//校验邮箱
func verifyEmail(email string) string {
	_email := ""
	flag := utils.VerifyEmailFormat(email)
	if flag == false {
		fmt.Println("请输入有效的邮箱！")
		fmt.Print("邮箱：")
		fmt.Scanln(&_email)
		verifyEmail(_email)
	} else {
		_email = email
	}
	return _email
}

//添加客户
func (c *customerView) add() {
	fmt.Println("-----------------添加客户-----------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	sex := ""
	fmt.Scanln(&sex)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("手机号：")
	phone := ""
	fmt.Scanln(&phone)
	phone = verifyPhone(phone)
	fmt.Print("邮箱：")
	email := ""
	fmt.Scanln(&email)
	email = verifyEmail(email)
	customer := model.NewCustomer2(name, age, sex, phone, email)
	if c.customerService.Add(customer) {
		fmt.Println("-----------------添加完成-----------------")
	} else {
		fmt.Println("-----------------添加失败-----------------")
	}
}

//修改客户
func (c *customerView) edit() {
	fmt.Println("-----------------修改客户-----------------")
	fmt.Print("请选择删除客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	index := c.customerService.FindById(id)

	if -1 == index {
		fmt.Println("-----------------输入id不存在，请重新输入--------------")
	} else {
		customer := c.customerService.GetInfoById(id)
		fmt.Printf("姓名(%v)：", customer.Name)
		name := ""
		fmt.Scanln(&name)
		fmt.Printf("性别(%v)：", customer.Gender)
		sex := ""
		fmt.Scanln(&sex)
		fmt.Printf("年龄(%d)：", customer.Age)
		age := 0
		fmt.Scanln(&age)
		fmt.Printf("手机号(%v)：", customer.Phone)
		phone := ""
		fmt.Scanln(&phone)
		phone = verifyPhone(phone)
		fmt.Printf("邮箱(%v)：", customer.Email)
		email := ""
		fmt.Scanln(&email)
		email = verifyEmail(email)
		customerModel := model.NewCustomer2(name, age, sex, phone, email)
		if c.customerService.Update(id, customerModel) {
			fmt.Println("-----------------修改完成-----------------")
		} else {
			fmt.Println("-----------------修改失败-----------------")
		}
	}

}

//删除某个客户的信息
func (c *customerView) delete() {
	fmt.Println("-----------------删除客户-----------------")
	fmt.Print("请选择删除客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Print("是否确认删除（Y/N）：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			if !c.customerService.Delete(id) {
				fmt.Println("你要删除的ID编号不存在！")
				break
			} else {
				fmt.Println("-----------------删除完成-----------------")
				break
			}
		} else if choice == "N" || choice == "n" {
			break
		} else {
			fmt.Print("您的输入有误！请重新输入（Y/N）：")
		}
	}
}

//退出
func (c *customerView) exit() {
	fmt.Print("是否确认退出系统？（Y/N）：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			c.loop = true
			break
		} else if choice == "N" || choice == "n" {
			break
		} else {
			fmt.Print("您的输入有误！请重新输入（Y/N）：")
		}
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
