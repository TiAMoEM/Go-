package main

import (
	"fmt"
)

type HeroNode struct {
	No       int
	name     string
	nickname string
	pre 	 *HeroNode
	next     *HeroNode
}

//方法一：在链表的最后插入数据
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//方法二：根据No的编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.No >= newHeroNode.No {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Printf("sorry, No=%d is exist\n", newHeroNode.No)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

	temp.next = newHeroNode
	newHeroNode.pre = temp
}

func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.No == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("sorry, the id doesn't exist")
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("It's empyt......")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.next.No, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		No:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		No:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		No:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	//hero4 := &HeroNode{
	//	No:       4,
	//	name:     "吴用",
	//	nickname: "智多星",
	//}

	InsertHeroNode(head, hero1)
	InsertHeroNode2(head,hero2)
	InsertHeroNode(head, hero3)
	ListHeroNode(head)
	DelHeroNode(head, 2)
	ListHeroNode(head)

}
