package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan's value is %v, intChan's address is %p\n", intChan, &intChan)

	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	fmt.Printf("channel len = %v, cap = %v\n", len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan
	fmt.Println("num2 =", num2)
	fmt.Printf("channel len = %v, cap = %v\n", len(intChan), cap(intChan))

	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3 =", num3, "num4 =", num4)
}
