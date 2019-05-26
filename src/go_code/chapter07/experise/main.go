package main

import "fmt"

func main() {
	var n1 int = 10
	var n2 int = 20
	var max int
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("max =", max)
}
