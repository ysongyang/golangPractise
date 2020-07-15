package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int //指向队列的前端(不含首)
	real    int //指向队列的尾部
}

func (queue *Queue) addQueue(val int) (err error) {
	if queue.real == queue.maxSize-1 {
		return errors.New("queue full")
	}
	queue.real++ //后移
	//fmt.Printf("real = %d\n", queue.real)
	queue.array[queue.real] = val
	return
}

func (queue *Queue) getQueue() (val int, err error) {
	if queue.real == queue.front {
		return -1, errors.New("queue empty")
	}
	queue.front++ //取一个数据  往后移一位
	val = queue.array[queue.front]
	return val, nil
}

func (queue *Queue) showQueue() {
	fmt.Println("Queue队列当前的情况是：")
	for i := queue.front + 1; i <= queue.real; i++ {
		fmt.Printf("array[%d] = %v\n", i, queue.array[i])
	}
}

func main() {
	queue := &Queue{maxSize: 5, front: -1, real: -1}
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
