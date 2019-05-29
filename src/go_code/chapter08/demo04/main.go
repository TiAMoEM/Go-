package main

import (
	"fmt"
)

func main() {
	var str string = "hello, world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c, \n", str[i])
	}

	fmt.Println()

	var str1 string = "abc, ok北京"
	for index, val := range str1 {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

}
