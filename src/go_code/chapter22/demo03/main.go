package main

import "fmt"
import "unicode/utf8"

func FizzBuzz() {
	i := 1
	for i < 101 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
		i++
	}
}

func ReverseString(str string) string {
	s := []rune(str)
	// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// }
	// return string(s)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return string(s)

}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
	FizzBuzz()

	str := "dsjkaisdyfhaisdhf...i a suhfi"
	fmt.Printf("String %s\n Length: %d, Runes: %d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
	s := str
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", s)
	fmt.Printf("After: %s\n", string(r))
	str1 := "foobar"
	fmt.Println(ReverseString(str1))
}
