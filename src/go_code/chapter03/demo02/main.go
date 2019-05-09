package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 = 100
	fmt.Printf("n1的类型为: %T \n", n1)
	fmt.Printf("n1占用的字节数为: %d ", unsafe.Sizeof(n1))
}
