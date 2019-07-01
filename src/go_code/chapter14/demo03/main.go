package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名：%v，年龄：%v，成绩：%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

//Golang中的继承为如下写法
//在结构体中嵌入匿名结构体
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生~~~~~")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生~~~~~")
}

func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "Bill"
	pupil.Student.Age = 24
	pupil.testing()
	pupil.Student.SetScore(90)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "Rose"
	graduate.Student.Age = 24
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()

}
