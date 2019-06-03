package main

import "fmt"

func main() {
	n := 10
	fmt.Println(Fib(n))
	fmt.Println(f(n))
}

//练习1
func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

//练习2
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("input error")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
