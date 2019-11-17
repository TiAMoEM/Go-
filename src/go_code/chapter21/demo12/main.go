package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int
	Top    int
	arr    [20]int
}

//入栈操作
func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("this stack is full")
		return errors.New("Stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

//出栈操作
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("this stack is empty")
		return 0, errors.New("Stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

//遍历栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("this stack is empty")
		return
	}
	fmt.Println("栈的情况如下...")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d\n", i, this.arr[i])
	}
}

func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num1 * num2
	case 43:
		res = num1 + num2
	case 45:
		res = num1 - num2
	case 47:
		res = num1 / num2
	default:
		fmt.Println("运算符错误")
	}
	return res
}

func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	example := "30+30*6-4-6"
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		ch := example[index : index+1]
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) {
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else {
			keepNum += ch
			if index == len(example)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operStack.IsOper(int([]byte(example[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}
		if index+1 == len(example) {
			break
		}
		index++
	}
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		numStack.Push(result)
	}
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s = %v", example, res)
}
