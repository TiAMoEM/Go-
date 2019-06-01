package main

import (
	"go_code/chapter09/demo01/utils"
	// "go_code"
	"fmt"
)

func main() {
	var n1 = 1.2
	var n2 = 3.4
	var operator byte = '+'
	fmt.Println(utils.Cal(n1, n2, operator))
}
