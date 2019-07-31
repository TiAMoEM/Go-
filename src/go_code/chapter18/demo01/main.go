package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func reflectTest(b interface{}) {
	//获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType =", rTyp)
	//获取reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2 =", n2)
	fmt.Printf("rVal = %v, rVal type = %T\n", rVal, rVal)

	//将rVal转成interfa{}
	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2 =", num2)
}

func reflectTest02(b interface{}) {
	//获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType =", rTyp)
	//获取reflect.Value
	rVal := reflect.ValueOf(b)

	iV := rVal.Interface()
	fmt.Printf("iV = %v, iV type = %T\n", iV, iV)

	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name = %v\n", stu.Name)
	}
}

func main() {
	var num int = 100
	reflectTest(num)
	stu := Student{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu)
}
