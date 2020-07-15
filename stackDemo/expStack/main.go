package main

//运算符利用栈来做计算

import (
	"errors"
	"fmt"
	"strconv"
)

//栈的使用案例 ,使用数组来模拟

type Stack struct {
	maxTop int     //栈最大可存放的个数
	top    int     //栈顶
	arr    [20]int //数组模拟栈
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
		return -1, errors.New("stack empty")
	}
	val = s.arr[s.top]
	s.top--
	return val, nil
}

//判断字符是不是运算符[+,-,*,/]
func (s *Stack) isOperator(val int) bool {
	//根据 ascii码对照表 来判断 43=+,45=-,42=*,47=/
	if val == 43 || val == 45 || val == 42 || val == 47 {
		return true
	}
	return false
}

//运算
func (s *Stack) cal(num1, num2 int, operator int) int {
	switch operator {
	case 43: //+
		return num2 + num1
	case 45: // -
		return num2 - num1
	case 42: //*
		return num2 * num1
	case 47: // /
		return num2 / num1
	default:
		fmt.Println(operator, "运算符错误")
		return 0
	}
}

//返回运算符的优先级
func (s *Stack) priority(operator int) int {
	if operator == 42 || operator == 47 {
		return 1
	} else if operator == 43 || operator == 45 {
		return 0
	}
	return -1
}

func main() {
	//数栈
	numStock := &Stack{
		maxTop: 20,
		top:    -1,
	}
	//符号栈
	operStock := &Stack{
		maxTop: 20,
		top:    -1,
	}

	//运算表达式
	exp := "301+31*6-120"

	//索引  扫描exp表达式
	index := 0

	num1 := 0
	num2 := 0
	oper := 0
	res := 0

	keepNum := ""
	for {
		//处理多位数

		ch := exp[index : index+1] //扫描运算符表达式
		temp := int([]byte(ch)[0])
		fmt.Printf("扫描的ch = %v, []byte(ch)[0] = %v temp = %v \n", ch, []byte(ch)[0], temp)
		if operStock.isOperator(temp) {
			//如果是个空栈 直接入栈
			if operStock.top == -1 {
				operStock.push(temp)
			} else {
				//要判断运算符的优先级
				if operStock.priority(operStock.arr[operStock.top]) >= operStock.priority(temp) {
					num1, _ = numStock.pop()  //数值栈弹出
					num2, _ = numStock.pop()  //数值栈弹出
					oper, _ = operStock.pop() //符号栈弹出
					res = operStock.cal(num1, num2, oper)
					//计算结果入栈
					numStock.push(res)
					//运算符号入栈
					operStock.push(temp)
				} else {
					operStock.push(temp)
				}
			}

		} else {

			keepNum += ch

			//处理十位数以上的逻辑
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStock.push(int(val))
			} else {
				if operStock.isOperator(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStock.push(int(val))
					keepNum = ""
				}
			}
		}

		if (index + 1) == len(exp) {
			break
		}
		index++
	}
	//取值做运算
	for {
		if operStock.top == -1 {
			break
		}
		num1, _ = numStock.pop()  //数值栈弹出
		num2, _ = numStock.pop()  //数值栈弹出
		oper, _ = operStock.pop() //符号栈弹出
		res = operStock.cal(num1, num2, oper)
		//计算结果入栈
		numStock.push(res)
	}
	//如果算法没有问题的话，numStock栈里应该只有最后计算的结果
	res, _ = numStock.pop()
	fmt.Println("表达式", exp, "结果：", res)
}
