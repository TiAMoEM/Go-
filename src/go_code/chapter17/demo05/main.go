package main

import (
	"fmt"
)

var Cat struct {
	Name string
	Age  int
}

func main() {
	//channel 存放int类型
	var intChan chan int
	intchan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Printf("num1 = %v, num2 = %v, num3 = %v\n", num1, num2, num3)

	//channel存放结构体变量
	var catChan chan Cat
	catChan = make(chan Cat, 10)
	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "tom~", Age: 20}
	catChan <- cat1
	catChan <- cat2

}
