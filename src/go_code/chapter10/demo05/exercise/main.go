package main

import (
	"fmt"
)

func Fib(n int) []uint64 {
	fibSlice := make([]uint64, n)
	fibSlice[0] = 1
	fibSlice[1] = 1
	for i := 2; i < n; i++ {
		fibSlice[i] = fibSlice[i-2] + fibSlice[i-1]
	}
	return fibSlice
}

func main() {
	fibSlice := Fib(20)
	fmt.Println("fibSlice =", fibSlice)
}
