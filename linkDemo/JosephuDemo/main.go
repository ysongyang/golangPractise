package main

import (
	"fmt"
	"strconv"
	"time"
)

//约瑟夫案例

//Boy结构体
type Boy struct {
	no   int
	name string
	next *Boy
}

//单向环形链表的构造
//num 小孩的个数
//return *Boy 返回该环形的链表第一个小孩的指针
func addBoy(num int) *Boy {
	first := &Boy{}
	temp := &Boy{}
	if num < 1 {
		fmt.Println("构建环形链表失败")
		return first
	}
	for i := 1; i <= num; i++ {
		//构建第一个小孩
		boy := &Boy{
			no:   i,
			name: "boy " + strconv.Itoa(i),
		}
		//构造循环链表需要一个辅助指针
		if i == 1 {
			first = boy
			temp = boy
			temp.next = first
		} else {
			temp.next = boy
			temp = boy
			temp.next = first //构成环线链表
		}
	}
	return first
}

//显示单向环形链表
func listBoy(first *Boy) int {
	fmt.Println("环形链表的信息如下：")
	if first.next == nil {
		fmt.Println("空链表，无法显示")
		return -1
	}
	num := 1
	tempNode := first
	for {
		fmt.Printf("小孩的信息 %v 为=[%d , %s]--> \n", tempNode, tempNode.no, tempNode.name)
		if tempNode.next == first {
			break
		}
		//判断链表后面是否还有链表
		tempNode = tempNode.next
		num++
	}
	fmt.Println()
	return num
}

func playGame(first *Boy, startNo, countNum int) {

	if first.next == nil {
		fmt.Println("空链表，无法进行...")
		return
	}

	/*if startNo <= listBoy(first) {

	}*/
	tail := first
	for {
		if tail.next == first { //说明tail到了头部
			break
		}
		//tail 指向 first 最后一个位置
		tail = tail.next
	}
	//让 first 根据startNo 移动
	for i := 1; i <= startNo-1; i++ {
		//同时移动
		first = first.next
		tail = tail.next
	}
	//开始数 contNum ，删除 first指向的小孩
	for {

		for i := 1; i <= countNum-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("小孩编号为 %d 出圈\n", first.no)
		//删除first指向的节点
		first = first.next //first移动一位
		tail.next = first  //tail.next 指向first  删除垃圾节点
		if tail == first {
			//说明圈中只有1个小孩了
			break
		}
	}
	fmt.Printf("最后出圈的小孩编号为 %d \n", first.no)
}

func main() {
	//执行时间
	t := time.Now()
	first := addBoy(200)
	listBoy(first)
	playGame(first, 20, 31)
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
}
