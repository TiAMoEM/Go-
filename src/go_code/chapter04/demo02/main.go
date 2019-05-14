package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		num1 int     = 99
		num2 float64 = 23.456
		// b      bool    = true
		// mychar byte    = 'a'
		str string
	)

	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str = %q \n", str, str)

	//
	str = strconv.FormatFloat(num2, "f", 10, 64)
	fmt.Printf("str type %T str = %q \n", str, str)

}
