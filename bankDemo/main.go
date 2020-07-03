package main

import (
	"fmt"
	"golangPractise/bankDemo/model"
	"golangPractise/utils"
)

func main() {

	var account model.Account
	for true {
		sno := 0
		fmt.Println("请根据菜单进行选择：")
		fmt.Println("1：银行开户")
		fmt.Println("2：银行存款")
		fmt.Println("3：银行取款")
		fmt.Println("4：查询余额")
		fmt.Println("5：退出")
		fmt.Scanln(&sno)
		if sno == 5 {
			fmt.Println("退出成功！")
			break
		}
		switch sno {
		case 1:
			var name string
			var pass string
			var money float64
			fmt.Println("请输入您的姓名：")
			fmt.Scanln(&name)
			fmt.Println("请输入您的银行卡密码：")
			fmt.Scanln(&pass)
			fmt.Println("请输入您开户时存的余额：")
			fmt.Scanln(&money)
			no := utils.GenerateCode()
			account.CardNo = no
			account.Name = name
			account.Password = pass
			account.Balance = money
			fmt.Printf("开户成功:\n卡号：%v\n姓名：%v\n当前余额：%v\n",
				account.CardNo, account.Name, account.Balance)
		case 2:
			var pass string
			var money float64

			fmt.Println("请输入您的银行卡密码：")
			fmt.Scanln(&pass)
			fmt.Println("请输入您存款金额：")
			fmt.Scanln(&money)
			account.Deposit(money, pass)
		case 3:
			var pass string
			var money float64

			fmt.Println("请输入您的银行卡密码：")
			fmt.Scanln(&pass)
			fmt.Println("请输入您取款金额：")
			fmt.Scanln(&money)
			account.Withdraw(money, pass)
		case 4:
			var pass string
			fmt.Println("请输入您的银行卡密码：")
			fmt.Scanln(&pass)
			account.Query(pass)
		}
	}
}
