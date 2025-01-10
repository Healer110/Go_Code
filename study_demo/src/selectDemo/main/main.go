package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 使用select解决从管道读取数据的阻塞问题

	// 定义一个管道，10个数据int
	var intChan chan int = make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 在定义一个管道，5个字符串

	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + strconv.Itoa(i)
	}
	// 传统方法，管道不关闭，遍历时会导致死锁 dead lock
	// 实际开发中，不能确定什么时候关闭管道，可以使用select的方式解决死锁问题
	// label:
	for {
		select {
		// 使用select获取数据，如果intChan没有关闭，不会阻塞而deadlock
		// 会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取数据：%d\n", v)
		case v := <-strChan:
			fmt.Printf("从strChan读取数据：%s\n", v)
		default:
			fmt.Println("无法获取数据...")
			// break label
			return
		}
	}
}
