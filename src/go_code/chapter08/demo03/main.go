package main

import (
	"fmt"
	_ "math"
)

func main() {
	// var key byte
	// fmt.Println("Please input a byte a, b, c, d, e, f, g")
	// fmt.Scanf("%c", &key)

	// switch key {
	// case 'a':
	// 	fmt.Println("Monday")
	// case 'b':
	// 	fmt.Println("Tuesday")
	// case 'c':
	// 	fmt.Println("Wednesday")
	// case 'd':
	// 	fmt.Println("Thursday")
	// case 'e':
	// 	fmt.Println("Friday")
	// case 'f':
	// 	fmt.Println("Saturday")
	// case 'g':
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("input error")
	// }

	//switch 穿透
	// var num int = 10

	// switch num {
	// case 10:
	// 	fmt.Println("11111")
	// 	fallthrough
	// case 20:
	// 	fmt.Println("22222")
	// 	fallthrough
	// case 30:
	// 	fmt.Println("33333")
	// default:
	// 	fmt.Println("44444")
	// }

	//Type switch
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Println("x 的类型是%T", i)
	case int:
		fmt.Println("x is int")
	case float64:
		fmt.Println("x is float64")
	case func(int) float64:
		fmt.Println("x is func(int)")
	default:
		fmt.Println("others")
	}

}
