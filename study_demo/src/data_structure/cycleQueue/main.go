package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int // 指向队首
	tail    int // 指向队尾
}

func (cq *CircleQueue) Push(val int) (err error) {
	if cq.IsFull() {
		fmt.Println("queue full")
		return errors.New("queue full")
	}
	cq.array[cq.tail] = val
	cq.tail = (cq.tail + 1) % cq.maxSize
	return

}

func (cq *CircleQueue) Pop() (val int, err error) {
	if cq.IsEmpty() {
		fmt.Println("queue enpty")
		return -1, errors.New("queue empty")
	}

	val = cq.array[cq.head]
	cq.head = (cq.head + 1) % cq.maxSize

	return
}

// 显示队列
func (cq *CircleQueue) ListQueue() {
	// 取出当前有多少个元素
	size := cq.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}
	fmt.Println("环形队列的值：")
	// 辅助变量
	tempHead := cq.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d] = %d \n", tempHead, cq.array[tempHead])
		tempHead = (tempHead + 1) % cq.maxSize
	}

}

// 判断环形队列是否已满
func (cq *CircleQueue) IsFull() bool {
	return (cq.tail+1)%cq.maxSize == cq.head
}

// 判断环形队列是否为空
func (cq *CircleQueue) IsEmpty() bool {
	return cq.tail == cq.head
}

// 环形队列有多少个元素
func (cq *CircleQueue) Size() int {
	return (cq.tail + cq.maxSize - cq.head) % cq.maxSize
}

func main() {
	cq := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	for i := 0; i < 5; i++ {
		cq.Push(i + 100)
	}

	cq.ListQueue()

}
