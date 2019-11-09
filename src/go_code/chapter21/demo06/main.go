package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode, "加入到环形列表")
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	temp.next = newCatNode
	newCatNode.next = head
}

func listCircleLink(head *CatNode) {
	fmt.Println("环形列表情况如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("It's an empty Linklist")
		return
	}

	for {
		fmt.Printf("猫的信息为[id = %d, name = %s]->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("It's an empty Linklist")
		return head
	}

	if temp.next == head {
		if temp.no == id {
			temp.next = nil
		}
		return head
	}

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		if temp.next == head {
			break
		}
		if temp.no == id {
			if temp == head {
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("猫猫 = %d\n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}

	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫 = %d\n", id)
		} else {
			fmt.Printf("对不起，没有no = %d\n", id)
		}
	}
	return head
}

func main() {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "Tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "Jerry",
	}

	cat3 := &CatNode{
		no:   3,
		name: "Ben",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	listCircleLink(head)

	head = DelCatNode(head, 3)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	listCircleLink(head)

}
