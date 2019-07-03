package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树...")
}

type LittleMonkey struct {
	Monkey
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "学会飞翔...")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "学会游泳...")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "TOM",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
