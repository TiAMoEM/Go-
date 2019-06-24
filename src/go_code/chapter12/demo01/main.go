package main

import (
	"fmt"
)

func main() {
	//map的声明
	var mv map[int]int
	//map声明后需要先make，make来给map分配数据空间
	mv = make(map[int]int, 10)
	//key值不能重复
	mv[1] = 10
	mv[2] = 20
	mv[3] = 10
	fmt.Println(mv)

	//map使用方式一如上
	//方式二
	cities := make(map[string]string)
	cities["no1"] = "Beijing"
	cities["no2"] = "ShangHai"
	cities["no3"] = "Hangzhou"
	fmt.Println(cities)

	//方式三
	heroes := map[string]string{
		"hero1": "Captain America",
		"hero2": "Iron Man",
		"hero3": "SpiderMan",
	}
	fmt.Println(heroes)
	fmt.Println()
	practise()

}

func practise() {
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["sex"] = "male"
	studentMap["stu01"]["age"] = "18"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Mary"
	studentMap["stu02"]["sex"] = "female"
	studentMap["stu02"]["age"] = "16"
	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu02"]["sex"])
}
