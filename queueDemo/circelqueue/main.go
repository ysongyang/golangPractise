package main

//环形链表队列案例

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int //头部 0
	tail    int //尾部 0
}

//入队列
func (cq *CircleQueue) addQueue(val int) (err error) {
	if cq.isFullQueue() {
		return errors.New("queue full")
	}
	//先放值
	cq.array[cq.tail] = val
	//尾部移动
	cq.tail = (cq.tail + 1) % cq.maxSize
	return
}

//出队列
func (cq *CircleQueue) getQueue() (val int, err error) {
	if cq.isEmptyQueue() {
		return -1, errors.New("queue empty")
	}
	//取一个数据   指向队首 包含首元素
	val = cq.array[cq.head]
	//往后移一位
	cq.head = (cq.head + 1) % cq.maxSize

	return val, nil
}

//显示队列
func (cq *CircleQueue) showQueue() {
	//去除当前队列有多少个元素
	size := cq.sizeQueue()
	if size == 0 {
		fmt.Println("queue empty")
		return
	}
	tempHead := cq.head
	fmt.Println("Queue队列当前的情况是：")
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d] = %v\n", tempHead, cq.array[tempHead])
		tempHead = (tempHead + 1) % cq.maxSize
	}
}

//判断队列是否满
func (cq *CircleQueue) isFullQueue() bool {
	return (cq.tail+1)%cq.maxSize == cq.head
}

//判断队列是否为空
func (cq *CircleQueue) isEmptyQueue() bool {
	return cq.tail == cq.head
}

//取出环形队列的元素
func (cq *CircleQueue) sizeQueue() int {
	//关键算法
	return (cq.tail + cq.maxSize - cq.head) % cq.maxSize
}

func main() {
	queue := &CircleQueue{maxSize: 5, head: 0, tail: 0}
	var key string
	var val int
	for {
		fmt.Println("1、输入add 添加数据到队列")
		fmt.Println("2、输入get 从队列获取数据")
		fmt.Println("3、输入show 显示队列数据")
		fmt.Println("4、输入exit 退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入要入队列的数：")
			fmt.Scanln(&val)
			if err := queue.addQueue(val); err != nil {
				fmt.Println("入队列失败", err.Error())
			} else {
				fmt.Println("入队列成功 ^_^")
			}
			fmt.Println()
		case "get":
			val, err := queue.getQueue()
			if err != nil {
				fmt.Println("获取队列失败", err.Error())
			} else {
				fmt.Println("获取队列成功：", val)
			}
			fmt.Println()
		case "show":
			queue.showQueue()
			fmt.Println()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("无效的命令")
		}
	}
}
