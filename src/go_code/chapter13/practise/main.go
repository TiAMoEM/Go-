package main

import (
	"fmt"
)

type MethodUtils struct {
}

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (mu MethodUtils) Print() {
	for i := 1; i < 11; i++ {
		for j := 1; j < 9; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) Area(lengh float64, width float64) float64 {
	return lengh * width
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

func (c *Calculator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = c.Num1 + c.Num2
	case '-':
		res = c.Num1 - c.Num2
	case '*':
		res = c.Num1 * c.Num2
	case '/':
		res = c.Num1 / c.Num2
	default:
		fmt.Println("operator is wrong")
	}
	return res
}

func main() {
	var mu MethodUtils
	var c Calculator
	c.Num1 = 20
	c.Num2 = 10
	fmt.Println(c.getRes('+'))
	mu.Print2(3, 4)
}
