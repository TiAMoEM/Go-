package main

//需要在环境变量下的src包下面做工程
import (
	"fmt"
	"go_code/chapter09/demo01/utils"
)

func main() {
	var n1 = 1.2
	var n2 = 3.4
	var operator byte = '+'
	fmt.Println(utils.Cal(n1, n2, operator))
}
