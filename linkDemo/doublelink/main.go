package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//双向链表案例

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //指向前一个节点
	next     *HeroNode //指向下一个节点
}

//给双向链表插入节点
func insertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//先找到当前链表最后节点
	tempNode := head
	//for循环寻找最后一个节点
	for {
		if tempNode.next == nil {
			break
		}
		tempNode = tempNode.next //让 tempNode 不断指向下一个节点
	}
	tempNode.next = newHeroNode
	newHeroNode.pre = tempNode
}

//根据编号给双向链表插入节点
func insertHeroNodeSort(head *HeroNode, newHeroNode *HeroNode) {
	//先找到当前链表适当的节点
	tempNode := head
	flag := true
	//for循环寻找最后一个节点
	for {
		if tempNode.next == nil {
			break
		} else if tempNode.next.no <= newHeroNode.no { //  < 倒序   > 升序
			break
		} else if tempNode.next.no == newHeroNode.no {
			flag = false
			break
		}
		tempNode = tempNode.next //让 tempNode 不断指向下一个节点
	}
	if !flag {
		fmt.Println("已经存在当前编号", newHeroNode.no)
		return
	} else {
		newHeroNode.next = tempNode.next

		newHeroNode.pre = tempNode
		if tempNode.next != nil {
			tempNode.next.pre = newHeroNode
		}
		tempNode.next = newHeroNode

	}
}

//显示链表的所有节点
func listHeroNode(head *HeroNode) {
	fmt.Println("顺序打印：")
	tempNode := head
	if tempNode.next == nil {
		fmt.Println("空链表，无法显示")
		return
	}
	for {
		fmt.Printf("[%d , %s , %s]next--->\n", tempNode.next.no, tempNode.next.name, tempNode.next.nickname)
		//判断链表后面是否还有链表
		tempNode = tempNode.next
		if tempNode.next == nil {
			break
		}
	}
}

//显示链表的所有节点【逆序打印】
func listHeroNode2(head *HeroNode) {
	fmt.Println("逆序打印：")
	tempNode := head

	if tempNode.next == nil {
		fmt.Println("空链表，无法显示")
		return
	}
	//让 tempNode 定位到双向链表的最后节点
	for {
		if tempNode.next == nil {
			break
		}
		tempNode = tempNode.next
	}
	for {
		fmt.Printf("[%d , %s , %s]next--->\n", tempNode.no, tempNode.name, tempNode.nickname)
		//判断链表是不是头部
		//往前移动进行打印
		tempNode = tempNode.pre
		if tempNode.pre == nil {
			break
		}
	}
}

//删除一个节点
func delHeroNode(head *HeroNode, no int) {
	tempNode := head
	flag := false
	//for循环寻找最后一个节点   找到要删除的节点
	for {
		if tempNode.next == nil { //说明到链表最后
			break
		} else if tempNode.next.no == no {
			flag = true
			break
		}
		tempNode = tempNode.next //让 tempNode 不断指向下一个节点
	}
	if flag {
		tempNode.next = tempNode.next.next
		if tempNode.next != nil {
			tempNode.next.pre = tempNode
		}
		fmt.Println("已经找到当前编号", no)
		return
	} else {
		fmt.Println("没有找到这个编号。。。", no)
	}
}

func main() {
	//先创建一个头节点
	head := &HeroNode{}

	//执行时间
	t := time.Now()
	rand.Seed(time.Now().Unix())
	//创建一个1000个HeroNode节点
	for i := 0; i < 1000; i++ {
		heads := &HeroNode{
			no:       rand.Intn(10000),
			name:     "张伟伟" + strconv.Itoa(i),
			nickname: "昵称称呼" + strconv.Itoa(i),
		}
		insertHeroNodeSort(head, heads)
	}
	listHeroNode(head)
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
}
