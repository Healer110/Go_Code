package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int // 指向队列首
	rear    int // 指向队列尾
}

// 添加数据到队列
func (q *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if q.rear == q.maxSize-1 {
		fmt.Println("queue full")
		return errors.New("queue full")
	}

	q.rear++ // rear后移
	q.array[q.rear] = val
	return
}

// 从队列中取出数据
func (q *Queue) GetQueue() (val int, err error) {
	// 先判断是否为空
	if q.front == q.rear {
		fmt.Println("Queue empty")
		return -1, errors.New("queue empty")
	}
	q.front++
	val = q.array[q.front]
	return

}

// 显示队列，找到队首，遍历到队尾
func (q *Queue) ShowQeue() {
	// front指向队首，不包含队首
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, q.array[i])
	}
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	queue.AddQueue(10)
	queue.AddQueue(20)
	queue.AddQueue(30)
	queue.AddQueue(40)
	queue.AddQueue(50)
	queue.AddQueue(50)
	queue.ShowQeue()
	for i := 0; i < 6; i++ {
		val, err := queue.GetQueue()
		if err != nil {
			return
		}
		fmt.Println(val)
		queue.ShowQeue()
	}

}
