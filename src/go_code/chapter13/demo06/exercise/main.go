package main

import (
	"fmt"
)

type Dog struct {
	name   string
	age    int
	weight int
}

func (d *Dog) say() string {
	infoStr := fmt.Sprintf("Dog的信息 name = [%v], age = [%v], weight = [%v]",
		d.name, d.age, d.weight)
	return infoStr
}

func main() {
	var d = Dog{
		name:   "mimi",
		age:    3,
		weight: 10,
	}
	fmt.Println(d.say())
}
