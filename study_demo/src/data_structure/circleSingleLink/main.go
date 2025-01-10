package main

import (
	"fmt"
)

// 定义猫的结构体节点
type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 判断是不是添加第一个item
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 形成一个环形
		fmt.Println(newCatNode, "已经加入")
		return
	}

	// 先定义临时的变量
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入
	temp.next = newCatNode
	newCatNode.next = head

}

func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表：")
	fmt.Printf("head--> %p\n", head)
	temp := head
	if temp.next == nil {
		fmt.Println("Link Empty...")
		return
	}
	for {
		// fmt.Printf("[id = %d, name = %s]\n", temp.next.no, temp.next.name)
		fmt.Println(temp)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 删除一只猫
func DeleteCatNode(head *CatNode, id int) *CatNode {
	// 两个辅助，temp主管移动找匹配元素，helper做删除动作
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("这是一个空环形链表，无法删除...")
		return head
	}
	// 只有一个节点时，删除掉自己，这里将节点的next指向nil即可
	if temp.next == head {
		temp.next = nil
		return head
	}

	// 将helper定位到环形链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	// 节点数≥2个，删除的节点正好是head节点时
	flag := true
	for {
		if temp.next == head { // 找到最后一个item,还没有比较
			break
		}
		if temp.no == id {
			// 删除的元素是head时，将head顺移到下一位，最后要将新的head返回
			if temp == head {
				head = head.next
			}

			helper.next = temp.next
			// fmt.Printf("ID = %d 删除成功.\n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}

	// 上面没有删除
	if flag {
		// 匹配到最后一个元素
		if temp.no == id {
			helper.next = temp.next
		} else {
			fmt.Println("ID 未找到", id)
			return head
		}
	}
	fmt.Printf("ID = %d 删除成功.\n", id)
	return head

}

func main() {
	// 初始化一个头节点
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "Tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "ss",
	}

	cat3 := &CatNode{
		no:   3,
		name: "ee",
	}

	cat4 := &CatNode{
		no:   4,
		name: "zz",
	}

	cat5 := &CatNode{
		no:   5,
		name: "tt",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	InsertCatNode(head, cat4)
	InsertCatNode(head, cat5)
	ListCircleLink(head)
	DeleteCatNode(head, 2)
	ListCircleLink(head)

}
