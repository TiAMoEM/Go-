package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	//关闭后不能再写入内容进入到channel
	fmt.Println("success")
	n1 := <-intChan
	fmt.Println("n1 =", n1)

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	//channel的遍历需要先关闭
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v =", v)
	}

}
