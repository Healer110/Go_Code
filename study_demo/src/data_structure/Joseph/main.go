package main

import "fmt"

type Boy struct {
	No   int
	next *Boy // 指向下一个人的变量
}

// 编写一个函数，组建一个单项环形列表
// num 人的个数，返回该环形链表第一个人的指针地址
func AddBoy(num int) *Boy {
	first := &Boy{}
	currBoy := &Boy{}

	if num < 1 {
		fmt.Println("无法构建....num >=1")
		return first
	}

	// 构建
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		// 构成循环链表，需要一个辅助指针
		// 因为第一个人比较特殊
		if i == 1 {
			first = boy
			currBoy = boy
			currBoy.next = first
		} else {
			currBoy.next = boy
			currBoy = boy
			boy.next = first
		}
	}
	return first
}

// 显示单项环形列表
func ShowBoy(first *Boy) {
	if first.next == nil {
		fmt.Println("链表为空...")
		return
	}

	// 创建辅助指针,说明至少有一个元素
	fmt.Printf("first boy address: %p -- > ", first)
	currBoy := first
	for {
		fmt.Printf("boyID=%d [%p]-->", currBoy.No, currBoy.next)
		if currBoy.next == first {
			break
		}
		currBoy = currBoy.next
	}
	fmt.Println()
}

/*
	设置编号为1,2,3..n的n个人围坐一圈，约定编号为k(1<=K<=n)
	的人从1开始报数，数到m的那个人出列，他的下一个人又从1 开始数，
	数到m又出列，以此类推
*/
func PlayGame(first *Boy, startNo int, countNum int) {
	if first.next == nil {
		fmt.Println("链表为空...")
		return
	}

	// startNo 要在1~n之间
	tail := first
	var totalNum = 0
	for {
		totalNum++
		if tail.next == first {
			break
		}

		tail = tail.next
	}

	if startNo < 1 || startNo > totalNum {
		fmt.Printf("startNo要在%d ~ %d之间\n", 1, totalNum)
	}

	// 定义一个辅助节点，负责删除元素,指向最后一个元素
	// 这里使用上面的tail即可
	// 让first移动到startNo, 这样从startNo开始数，数countNum下，后进行删除即可
	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	// 继续移动，数countNum，然后删除first指向的节点,自己首先占用一次
	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("编号为%d的人出列...\n", first.No)
		// 删除first指向的节点
		first = first.next
		tail.next = first
		if tail == first {
			// 链表中只有一个人
			break
		}
	}
	fmt.Printf("最后出列的人的编号为%d...\n", first.No)
}

func main() {
	first := AddBoy(50)
	ShowBoy(first)
	PlayGame(first, 20, 31)
}
