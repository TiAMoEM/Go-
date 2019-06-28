package main

import (
	"fmt"
)

type Person struct {
	Name string
}

//给Person类型绑定一个方法
//test方法只能通过Person类型的变量来调用，也不能用其他类型变量来调用
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func (p Person) speak() {
	fmt.Println(p.Name, "is a good man")
}

func (p Person) calculate() {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Println(p.Name, "result =", res)
}

func (p Person) calculate2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "result2 =", res)
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var p Person
	p.Name = "Bill"
	p.test()
	p.speak()
	p.calculate()
	p.calculate2(20)
	fmt.Println(p.getSum(10, 20))
}
