package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil { //找到链尾
			break
		}
		temp = temp.next //temp不断指向下一个节点
	}
	temp.next = newHeroNode
}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//找到适当节点
	//辅助节点
	temp := head
	flag := true
	for {
		if temp.next == nil { //找到链尾
			break
		} else if temp.next.no >= newHeroNode.no {
			//newHeroNode应该插到temp节点后面，所以要退出循环 进行插入操作
			break
		} else if temp.next.no == newHeroNode.no {
			//说明我们链表中已经有no，就不能插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起 已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

//显示链表所有结点信息
func ListHeroNode(head *HeroNode) {
	temp := head //辅助节点
	//判断是否是空链表
	if temp.next == nil {
		fmt.Println("空表")
		return
	}
	//非空那就遍历链表
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func DelHeroNode(head *HeroNode, id int) {
	temp := head

	for {
		if temp.next == nil {
			fmt.Println("没炸到")
			break
		} else if temp.next.no == id {
			temp.next = temp.next.next
			break
		}
		temp = temp.next
	}
}
func main() {
	//先创建一个头节点
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "啦啦啦",
	}
	hero3 := &HeroNode{
		no:       2,
		name:     "林冲",
		nickname: "豹子头",
	}
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)
	ListHeroNode(head)
	DelHeroNode(head, 2)
	DelHeroNode(head, 3)
	DelHeroNode(head, 2)
	ListHeroNode(head)
}
