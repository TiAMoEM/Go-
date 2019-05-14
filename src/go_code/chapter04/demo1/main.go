package main

import "fmt"

//使用fmt中的Sprintf函数
func main() {
	var (
		num1   int     = 99
		num2   float64 = 23.456
		b      bool    = true
		mychar byte    = 'a'
		str    string
	)

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str = %q \n", str, str)

}
