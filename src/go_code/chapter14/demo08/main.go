package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	//类型断言，判断a是否指向Point类型的变量，如果是就转成Point类型并赋给b变量，否则报错
	b = a.(Point)
	fmt.Println(b)

	var x interface{}
	var b2 float32 = 1.1
	x = b2
	// y := x.(float32)
	// fmt.Printf("y的类型是%T,值是%v\n", y, y)

	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T,值是%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}

}
