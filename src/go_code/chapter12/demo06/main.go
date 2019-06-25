package main

import "fmt"

func main() {
	map1 := make(map[int]int, 3)
	map1[10] = 100
	map1[1] = 10
	map1[4] = 40
	map1[8] = 80
	fmt.Println(map1)
	modify(map1)
	fmt.Println(map1)

	//map的容量达到后，再增加元素，会自动扩容

	//map的value经常使用struct类型

	students := make(map[string]Stu, 10)
	stu1 := Stu{"Tom", 18, "BeiJing"}
	stu2 := Stu{"Mary", 16, "ShangHai"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	for k, v := range students {
		fmt.Printf("标号是%v \n", k)
		fmt.Printf("名字是%v \n", v.Name)
		fmt.Printf("年龄是%v \n", v.Age)
		fmt.Printf("地址是%v \n", v.Address)
		fmt.Println()
	}
}

func modify(m map[int]int) {
	m[1] = 9999
}

type Stu struct {
	Name    string
	Age     int
	Address string
}
