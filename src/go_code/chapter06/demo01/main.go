package main

import "fmt"

func main() {
	var num int = 9
	fmt.Printf("num address = %v \n", &num)

	var ptr *int
	ptr = &num
	*ptr = 20
	fmt.Println("num =", num)
}
