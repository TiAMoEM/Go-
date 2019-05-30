package main

import (
	"fmt"
)

/*
  *
 ***
*****
*/

func main() {
	var level int = 5

	for i := 1; i <= level; i++ {
		for k := 1; k <= level-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == level {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
