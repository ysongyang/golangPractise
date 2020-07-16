package main

import (
	"fmt"
	"os"
)

//hashTable使用案例

type Emp struct {
	Id   int
	Name string
	Next *Emp //下一个节点
}

func (emp *Emp) show() {
	fmt.Printf("链表 %d 员工ID：%d 姓名：%s \n", emp.Id%7, emp.Id, emp.Name)
}

//初始化节点存放Emp
type EmpLink struct {
	Head *Emp
}

//添加员工的方法
//保证添加时 编号从小到大
func (el *EmpLink) addEmp(emp *Emp) {
	curEmp := el.Head //定义一个辅助指针
	//如果当前的EmpLink 是一个空链表
	if curEmp == nil {
		el.Head = emp //添加完成
		return
	}
	var preEmp *Emp = nil
	//如果不是一个空链表 ，给emp找到对应的位置并插入
	// 让 curEmp 和 emp 比较，让preEmp 保持在 curEmp 前面
	for {
		if curEmp != nil {
			//fmt.Println("curEmp.id = ", curEmp.Id) //16
			//fmt.Println("emp.id = ", emp.Id)       //9
			//如果 curEmp.ID 大于 emp.Id
			if curEmp.Id > emp.Id {
				//找到位置后
				break
			}
			preEmp = curEmp //保证同步

			//继续寻找下一个
			curEmp = curEmp.Next
		} else {
			break
		}
	}
	//fmt.Println("preEmp = ", preEmp)

	//说明插入的id 往前面
	//如果preEmp==nil 说明插入的Emp 在头部
	if preEmp == nil {
		el.Head = emp
		emp.Next = curEmp
	} else {
		//是否将emp添加到链表最后
		preEmp.Next = emp
		emp.Next = curEmp
	}

}

//显示当前链表的信息
func (el *EmpLink) showLink(id int) {
	if el.Head == nil {
		fmt.Println(id, "当前链表为空")
		return
	}

	curEmp := el.Head //辅助节点
	for {
		if curEmp != nil {
			fmt.Printf("链表 %d 员工ID：%d 姓名：%s -->", id, curEmp.Id, curEmp.Name)
			curEmp = curEmp.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (el *EmpLink) findById(id int) *Emp {
	if el.Head == nil {
		fmt.Println(id, "当前链表为空,未找到")
		return nil
	}
	curEmp := el.Head //辅助节点
	for {
		if curEmp != nil && curEmp.Id == id {
			return curEmp
		} else if curEmp == nil {
			break
		}
		curEmp = curEmp.Next
	}
	return nil
}

type HashTable struct {
	LinkArr [7]EmpLink
}

//添加的方法  参数要用指针，防止值拷贝
func (ht *HashTable) addEmp(emp *Emp) {
	//使用散列函数确定将该emp添加到哪一个链表
	linkNo := ht.hashFun(emp.Id)
	//使用对应的链表添加
	ht.LinkArr[linkNo].addEmp(emp)
}

//显示所有链表
func (ht *HashTable) listAll() {
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].showLink(i)
	}
}

func (ht *HashTable) findById(id int) *Emp {
	linkNo := ht.hashFun(id)
	return ht.LinkArr[linkNo].findById(id)
}

//散列方法
func (ht *HashTable) hashFun(id int) int {

	return id % 7
}

func main() {
	var key int
	var id int
	var name string
	ht := &HashTable{}
	for {
		fmt.Println("请选择雇员菜单：")
		fmt.Println("1,添加雇员")
		fmt.Println("2,显示雇员")
		fmt.Println("3,查找雇员")
		fmt.Println("0,退出系统")
		fmt.Println("请输入您的选择：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入雇员ID号：")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员名字：")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			ht.addEmp(emp)
		case 2:
			ht.listAll()
		case 3:
			fmt.Println("请输入雇员ID号：")
			fmt.Scanln(&id)
			emp := ht.findById(id)
			if emp != nil {
				emp.show()
			}
		case 0:
			os.Exit(0)
		default:
			fmt.Println("请输入有效的口令！")
		}
	}
}
