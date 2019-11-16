package main

import "fmt"

type Boy struct {
	No   int
	Next *Boy
}

func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}

		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

func showBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}

	curBoy := first
	for {
		fmt.Printf("小孩编号为%d ->", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func playGame(first *Boy, startNo int, countNum int) {
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}

	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	fmt.Println()

	for {
		for i := 1; i < countNum; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为 %d 出圈 \n", first.No)
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}

	fmt.Printf("小孩编号为 %d 出圈 \n", first.No)
}

func main() {
	first := AddBoy(500)
	showBoy(first)
	playGame(first, 20, 31)

}
