package model

import "fmt"

//声明一个Customer结构体，客户信息
type Customer struct {
	Id     int
	name   string
	age    int
	gender string
	phone  string
	email  string
}

//使用工厂模式，返回customer实例
func NewCustomer(id int, name string, age int, gender string, phone string, email string) Customer {
	return Customer{id, name, age, gender, phone, email}
}

func NewCustomer2(name string, age int, gender string, phone string, email string) Customer {
	return Customer{name: name, age: age, gender: gender, phone: phone, email: email}
}

//返回当前客户
func (c *Customer) GetInfo() string {
	info := fmt.Sprintf("%d\t%v\t%v\t%d\t%v\t%v", c.Id, c.name, c.gender, c.age, c.phone, c.email)
	return info
}
