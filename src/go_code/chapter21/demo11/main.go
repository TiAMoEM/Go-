package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int
	Top    int
	arr    [6]int
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

func main() {
	stack := &Stack{
		MaxTop: 10,
		Top:    -1,
	}
	stack.List()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.List()
}
