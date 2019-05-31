package main

import (
	"fmt"
	_ "math/rand"
	_ "time"
)

func main() {
	// var count int = 0
	// for {
	// 	rand.Seed(time.Now().UnixNano())
	// 	n := rand.Intn(100) + 1
	// 	fmt.Println("n=", n)
	// 	count++
	// 	if n == 99 {
	// 		break
	// 	}
	// }
	// fmt.Println("aaaaaa", count)

label2:
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label2
			}
			fmt.Println("j=", j)
		}
	}
}
