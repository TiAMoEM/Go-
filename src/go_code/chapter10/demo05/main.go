package main

import (
	"fmt"
)

func main() {
	var arr [5]int = [...]int{2, 4, 6, 8, 20}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\n", i, slice[i])
	}

	fmt.Println()

	for i, v := range slice {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}

	slice2 := slice[1:2]
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(arr) //2, 4, 6, 8, 20
	slice2[0] = 10
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(arr) //2, 4, 10, 8, 20
}
