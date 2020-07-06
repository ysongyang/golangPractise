package model

import (
	"fmt"
)

type Account struct {
	CardNo   string
	Name     string
	Password string
	Balance  float64
}

//存款
func (account *Account) Deposit(money float64, password string) bool {
	if (*account).Password != password {
		fmt.Println("密码输入有误！")
		return false
	}

	//存款金额判断
	if money <= 0 {
		fmt.Println("请输入有效的存款金额！")
		return false
	}

	(*account).Balance += money
	fmt.Println("存款成功，当前账户余额：", (*account).Balance)
	return true
}

//取款
func (account *Account) Withdraw(money float64, password string) bool {
	if (*account).Password != password {
		fmt.Println("密码输入有误！")
		return false
	}

	//存款金额判断
	if money <= 0 {
		fmt.Println("请输入有效的取款金额！")
		return false
	}

	if money > (*account).Balance {
		fmt.Println("您的取款金额大于您的账户余额！")
		return false
	}

	(*account).Balance -= money
	fmt.Println("取款成功，当前账户余额：", (*account).Balance)
	return true
}

//取款
func (account *Account) Query(password string) bool {
	if (*account).Password != password {
		fmt.Println("密码输入有误！")
		return false
	}
	fmt.Printf("您的账号为：%v \t当前账户余额：%v\n", (*account).CardNo, (*account).Balance)
	return true
}
