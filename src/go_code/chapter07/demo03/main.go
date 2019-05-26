package main

import "fmt"

func main() {
	var (
		name   string
		age    byte
		salary float32
		isPass bool
	)
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水：")
	fmt.Scanln(&salary)

	fmt.Println("请输入是否通过考试：")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是%v \n 年龄是%v \n 薪水是%v \n 是否通过考试%v \n", name, age, salary, isPass)

}
