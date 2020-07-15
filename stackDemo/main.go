package main

import (
	"errors"
	"fmt"
)

//栈的使用案例 ,使用数组来模拟

type Stack struct {
	maxTop int    //栈最大可存放的个数
	top    int    //栈顶
	arr    [5]int //数组模拟栈
}

//入栈
func (s *Stack) push(val int) (err error) {
	//先判断栈是否满
	if s.top+1 == s.maxTop {
		fmt.Println("stack full", val)
		return errors.New("stack full")
	}
	s.top++
	s.arr[s.top] = val
	return
}

//显示栈
//从栈定开始遍历
func (s *Stack) list() {
	if s.top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("Stack list :")
	//curTop := s.top
	for i := s.top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d \n", i, s.arr[i])
	}
}

//出栈
func (s *Stack) pop() (val int, err error) {
	if s.top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	val = s.arr[s.top]
	s.top--
	return val, nil
}

func main() {
	stack := &Stack{
		maxTop: 5,
		top:    -1,
	}
	stack.push(12)
	stack.push(120)
	stack.push(112)
	stack.push(132)
	stack.push(1132)
	stack.push(15132)
	stack.list()
	if val, err := stack.pop(); err == nil {
		fmt.Println("出栈：", val)
	}
	if val, err := stack.pop(); err == nil {
		fmt.Println("出栈：", val)

	}
	if val, err := stack.pop(); err == nil {
		fmt.Println("出栈：", val)

	}
	stack.push(15132)
	stack.list()
}
