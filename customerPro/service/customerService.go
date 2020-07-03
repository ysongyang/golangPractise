package service

import "golangPractise/customerPro/model"

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
	customer := model.NewCustomer(customerService.customerNum, "张三", 35, "男", "13895648891", "zhangsan1699@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户信息
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}