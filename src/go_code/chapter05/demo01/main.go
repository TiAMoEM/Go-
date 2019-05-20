package main

//string类型转基本数据

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		str string = "true"
		b   bool
	)
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T b = %v\n", b, b)

	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type is %T n1 = %v\n", n1, n1)
	fmt.Printf("n2 type is %T n2 = %v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type is %T f1 = %v\n", f1, f1)
}
