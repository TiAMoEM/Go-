package main

import (
	"fmt"
	"math"
)

func main() {
	var t bool = true
	if t == false {
		fmt.Println("a")
	} else if t {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}

	var (
		a float64 = 2.0
		b float64 = 4.0
		c float64 = 2.0
	)
	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v, x2=%v", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1 = %v", x1)
	} else {
		fmt.Println("无解~~~")
	}

}
