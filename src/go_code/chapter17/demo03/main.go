package main

import (
	"fmt"
	"sync"
	"time"
)

//会出现资源争夺，出现以下报错
//fatal error: concurrent map writes
// var myMap = make(map[int]int, 10)

//解决方式一：使用全局变量加锁同步改进程序
var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i < n+1; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {
	for i := 1; i < 201; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}
