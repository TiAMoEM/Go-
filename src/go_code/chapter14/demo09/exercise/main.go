package main

import (
	"fmt"
)

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型,值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型,值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型,值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型,值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是 string 类型,值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是 Student 类型,值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数是 *Student 类型,值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数类型不确定,值是%v\n", index, x)
		}
	}
}

type Student struct {
	name string
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "Bill"
	address := "BeiJing"
	n4 := 300
	n5 := true
	n6 := Student{"ABC"}
	n7 := &Student{"Tom"}
	TypeJudge(n1, n2, n3, name, address, n4, n5, n6, n7)
}
