package main

import (
	"fmt"
	"go_code/chapter14/demo02/model"
)

func main() {
	p := model.NewPerson("Smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, "age =", p.GetAge(), "sal =", p.GetSal())
}
