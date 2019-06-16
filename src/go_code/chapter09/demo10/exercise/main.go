package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 10000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	start := time.Now().Unix()
	fmt.Println(start)
	test()
	end := time.Now().Unix()
	fmt.Println(end)
	fmt.Printf("耗时为%v秒\n", end-start)
}
