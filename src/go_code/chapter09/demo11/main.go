package main

import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("Type is %T, num1 value is %v, address is %v \n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 100
	fmt.Printf("Type is %T, num2 value is %v, address is %v, pointer is %v\n", num2, num2, &num2, *num2)
}
