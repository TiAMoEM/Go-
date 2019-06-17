package main

import (
	"fmt"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}

	for i, v := range arr {
		fmt.Printf("i = %v, v = %v\n", i, v)
		fmt.Printf("arr[%d] = %v\n", i, arr[i])
	}

	for _, v := range arr {
		fmt.Printf("元素的值为%v\n", v)
	}

}
