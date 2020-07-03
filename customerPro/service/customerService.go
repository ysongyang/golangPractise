package service

import (
	"golangPractise/customerPro/model"
)

//完成对customer的操作，包括CURD
type CustomerService struct {
	//定义Customer切片
	customers []model.Customer
	//编号，表示当前切片含有多少个数据，以及作为新客户的编号（id+1）
	customerNum int
}

//显示客户列表
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	//创建一个初始化客户
	customer := model.NewCustomer(customerService.customerNum, "张三", 35, "男", "13895648891", "zhangsan1688@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户信息
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

//添加客户方法  这里定义了指针/引用的CustomerService，保证数据切片是同一个
func (cs *CustomerService) Add(customer model.Customer) bool {
	//分配id
	cs.customerNum++ //编号自增
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

//编辑客户
func (cs *CustomerService) Update(id int, customer model.Customer) bool {
	index := cs.FindById(id)
	if -1 == index {
		return false
	}
	cs.customers[index].Name = customer.Name
	cs.customers[index].Age = customer.Age
	cs.customers[index].Gender = customer.Gender
	cs.customers[index].Phone = customer.Phone
	cs.customers[index].Email = customer.Email
	return true
}

//删除客户
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if -1 == index {
		return false
	}
	//删除切片中的数据
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

//返回当前id的客户信息
func (cs *CustomerService) GetInfoById(id int) model.Customer {
	return cs.customers[cs.FindById(id)]
}

//查找ID是否存在
func (cs *CustomerService) FindById(id int) int {
	index := -1
	for key, val := range cs.customers {
		if val.Id == id {
			index = key
		}
	}
	return index
}
