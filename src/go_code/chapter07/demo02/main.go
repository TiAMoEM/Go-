package main

import "fmt"

func main() {
	a := 100
	fmt.Println("a 的地址是", &a)

	var ptr *int = &a
	fmt.Println("ptr 指向的值是", *ptr)

}
