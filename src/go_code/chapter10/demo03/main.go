package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printChar() {
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}
}

func maxArr() {
	var intArr [6]int = [...]int{1, -1, 9, 90, 11, 9000}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal = %v, maxValIndex = %v\n", maxVal, maxValIndex)
}

func averageArr() {
	var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	fmt.Printf("sum = %v, average = %v\n", sum, float64(sum)/float64(len(intArr2)))
}

func reverseArr() {
	var intArr3 [6]int
	len := len(intArr3)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println("交换前~~~~~~", intArr3)
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr3[len-1-i]
		intArr3[len-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后~~~~~~", intArr3)
}

func main() {
	printChar()
	fmt.Println()
	maxArr()
	averageArr()
	reverseArr()

}
