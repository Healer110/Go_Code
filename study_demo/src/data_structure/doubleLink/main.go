package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode // 指向前一个节点
	next     *HeroNode // 表示指向下一个节点
}

// 给双向链表插入一个节点
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 先找到该链表的最后的节点
	// 创建一个辅助节点
	temp := head
	for {
		if temp.next == nil { // 表示找到了最后一个Node
			break
		}
		temp = temp.next // 未找到，check下一个的情况
	}
	// 将新节点加入到链表之后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

// 给双向列表插入一个新节点
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 先找到该链表适当的节点，进行插入
	temp := head
	flag := true
	for {
		if temp.next == nil {
			// 说明到链表的最后了
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("英雄排名已存在, no =", newHeroNode.no)
		return
	}
	// 将新节点加入到链表之后
	newHeroNode.next = temp.next
	newHeroNode.pre = temp
	if temp.next != nil {
		temp.next.pre = newHeroNode
	}
	temp.next = newHeroNode
}

// 双向链表删除节点
func DeleteHero(head *HeroNode, id int) {
	temp := head
	flag := false

	for {
		if temp.next == nil {
			// 说明到链表的最后了
			break
		} else if temp.next.no == id {
			// 找到了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		if temp.next.next != nil {
			temp.next.next.pre = temp
		}
		temp.next = temp.next.next

	} else {
		fmt.Println("要删除的英雄ID不存在...")
	}
}

// 显示链表的所有节点信息
func ListHeroNode(head *HeroNode) {
	fmt.Println("双向列表，正向打印...")
	temp := head
	// 先判断该链表是不是空链表
	if temp.next == nil {
		fmt.Println("link empty...")
		return
	}

	for {
		fmt.Printf("[%d, %s, %s]==> \n", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

// 显示双向链表的所有节点信息，逆向显示
func ListHeroNode2(head *HeroNode) {
	fmt.Println("双向列表，逆向打印...")
	temp := head
	// 先判断该链表是不是空链表
	if temp.next == nil {
		fmt.Println("link empty...")
		return
	}

	for {
		temp = temp.next
		if temp.next == nil {
			// 到列表的最后item了
			break
		}
	}

	for {
		fmt.Printf("[%d, %s, %s]==> \n", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {
	// 创建一个头节点，头结点用来标识链表的头部，不存放任何数据
	head := &HeroNode{}

	// 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	hero4 := &HeroNode{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}

	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero4)

	ListHeroNode(head)
	// 逆序打印
	ListHeroNode2(head)

	DeleteHero(head, 3)
	ListHeroNode(head)

	// InsertHeroNode2(head, hero3)
	// InsertHeroNode2(head, hero4)
	// InsertHeroNode2(head, hero2)
	// InsertHeroNode2(head, hero1)
	// ListHeroNode(head)
	// DeleteHero(head, 6) // 不存在
	// DeleteHero(head, 3)
	// ListHeroNode(head)
}
