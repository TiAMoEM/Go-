package main

import (
	"fmt"
)

// 1) 输入整数，打印出对应的金字塔

func printPyramid(totalLevel int) {
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Println("please input Level number")
	fmt.Scanln(&n)
	printPyramid(n)
}
