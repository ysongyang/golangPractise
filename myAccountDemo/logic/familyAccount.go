package logic

import "fmt"

type FamilyAccount struct {
	//声明字段
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	details string
}

//构造方法
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0,
		note:    "",
		details: "",
	}
}

//方法

func (fa *FamilyAccount) ShowMenu() {
	for {
		fmt.Println("--------------------家庭收支记账软件--------------------")
		fmt.Println("                     1、收支明细                     ")
		fmt.Println("                     2、登记收入                     ")
		fmt.Println("                     3、登记支出                     ")
		fmt.Println("                     4、退出软件                     ")
		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&fa.key)
		switch (*fa).key {
		case "1":
			fa.showDetails()
		case "2":
			fa.income()
		case "3":
			fa.pay()
		case "4":
			fa.exit()
		}
		if !fa.loop {
			break
		}
	}
}

//显示当前收支明细记录
func (fa *FamilyAccount) showDetails() {
	fmt.Println("--------------------当前收支明细记录--------------------")
	if len(fa.details) > 0 {
		fmt.Println("明细\t账户金额\t收支金额\t说   明\n")
		fmt.Println(fa.details)
	} else {
		fmt.Println("当前没有收支明细记录，快来记录一笔吧！")
	}
}

//登记收入
func (fa *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&fa.money)
	fmt.Println("本次收入说明：")
	fmt.Scanln(&fa.note)
	fa.balance += fa.money
	fa.details += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n", fa.balance, fa.money, fa.note)
}

//登记支出
func (fa *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance {
		fmt.Println("您的账户余额不足!")
		return
	}
	fmt.Println("本次支出说明：")
	fmt.Scanln(&fa.note)
	fa.balance -= fa.money
	fa.details += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n", fa.balance, fa.money, fa.note)
}

//退出
func (fa *FamilyAccount) exit() {
	fmt.Println("您确定要退出吗？(y/n)")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("您的输入有误，请重新输入！(y/n)")
	}
	fmt.Println("退出成功...")

	if choice == "y" {
		fa.loop = false
	}
	return
}
