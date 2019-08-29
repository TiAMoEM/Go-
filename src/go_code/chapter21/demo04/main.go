package main

import (
	"fmt"
)

type HeroNode struct {
	No       int
	name     string
	nickname string
	next     *HeroNode
}

//方法一：在单链表的最后插入数据
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
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
		fmt.Println("已经存在No=", newHeroNode.No)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也....")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s] ==>", temp.next.No, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
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
	} else {
		fmt.Println("sorry, 要删除的id不存在")
	}

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
	hero4 := &HeroNode{
		No:       4,
		name:     "吴用",
		nickname: "智多星",
	}
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero4)
	ListHeroNode(head)
}
