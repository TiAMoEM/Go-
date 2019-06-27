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

func main() {
	var p Person
	p.Name = "Bill"
	p.test()
}
