package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//单链表案例
//做排名比较好用
//缺陷是查找从头往尾进行查找，不能自我删除

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //指向下一个节点
}

//给链表插入节点【普通的链表插入，没有使用价值】
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
}

//根据编号给链表插入节点【可以做排行榜】
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
		tempNode.next = newHeroNode
	}
}

//显示链表的所有节点
func listHeroNode(head *HeroNode) {
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
		fmt.Println("已经找到当前编号", no)
		return
	} else {
		fmt.Println("没有找到这个编号。。。", no)
	}
}

func main() {
	//执行时间
	t := time.Now()

	//先创建一个头节点
	head := &HeroNode{}
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
