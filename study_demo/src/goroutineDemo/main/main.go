package main

import (
	"fmt"
	"time"
)

// 在主线程中开启一个goroutine协程，每隔一秒打印一句话
// 主线程也每秒输出一句话，10次后退出程序
func main() {
	go demo01() // 开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("main function output ", i)
		time.Sleep(time.Second)
	}
}

func demo01() {
	for i := 0; i < 10; i++ {
		fmt.Println("demo01 output ", i)
		time.Sleep(time.Second)
	}
}
