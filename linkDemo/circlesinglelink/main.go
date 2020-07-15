package main

import "fmt"

//环形单向链表案例

type CatNode struct {
	no   int //编号
	name string
	next *CatNode
}

func insertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //环形指向
		fmt.Println("第一只猫", newCatNode, "加入到环形的链表")
		return
	}
	//定义一个临时节点变量
	tempNode := head
	for {
		if tempNode.next == head {
			break
		}
		tempNode = tempNode.next
	}
	//加入链表
	tempNode.next = newCatNode
	newCatNode.next = head
}

//显示环形链表的所有节点
func listCatNode(head *CatNode) {
	fmt.Println("环形链表的信息如下：")
	tempNode := head
	if tempNode.next == nil {
		fmt.Println("空链表，无法显示")
		return
	}
	for {
		fmt.Printf("猫的信息 %v 为=[%d , %s]--> \n", tempNode, tempNode.no, tempNode.name)
		if tempNode.next == head {
			break
		}
		//判断链表后面是否还有链表
		tempNode = tempNode.next
	}
}

//删除
func delCatNode(head *CatNode, no int) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表，无法删除")
		return
	}
	//如果只有一个节点
	if temp.next == head {
		temp.next = nil
		return
	}
	//如果有2个以上的节点
	helper := head

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := false
	for {
		if temp.next == head {
			break
		}
		if temp.no == no {
			if temp == head {
				head = head.next
			}
			helper.next = temp.next
			fmt.Println(no, "已被删除")
			flag = true
			break
		}
		temp = temp.next     //用于比较
		helper = helper.next //用于找到要删除的节点
	}
	if !flag {
		if temp.no == no {
			fmt.Println(no, "已被删除")
			helper.next = temp.next
		} else {
			fmt.Println(no, "没有找到")
		}
	}
}

func main() {
	head := &CatNode{} //初始化头节点

	cat1 := &CatNode{
		no:   1000,
		name: "波斯猫",
	}

	cat2 := &CatNode{
		no:   1001,
		name: "汤姆猫",
	}

	cat3 := &CatNode{
		no:   1002,
		name: "金丝猫",
	}

	insertCatNode(head, cat1)
	insertCatNode(head, cat2)
	insertCatNode(head, cat3)
	listCatNode(head)

	delCatNode(head, 1001)

	listCatNode(head)

}
