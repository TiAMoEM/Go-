package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 = 100
	var str = "hello, world"
	fmt.Printf("n1的类型为: %T \n", n1)
	fmt.Printf("n1占用的字节数为: %d \n", unsafe.Sizeof(n1))
	fmt.Println(str[0:2])

	var num1 int64 = 1000
	var num2 int8 = int8(num1)
	fmt.Println("num2 is ", num2)

}
