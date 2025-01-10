package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
开启一个writeData协程，向管道intChan中写入50个整数
开启一个readData协程，向管道initChan中读取writeData写入的数据
主线程需要等待writeData和readData协程都完成工作才能退出
*/

var intChan chan int = make(chan int, 50)
var exitChan chan bool = make(chan bool, 1)

func main() {
	go writeData()
	go readData()
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
		fmt.Println("主线程退出.....")
	}
}

// 写线程
func writeData() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 50; i++ {
		randNum := rand.Intn(100) + 1
		intChan <- randNum
		fmt.Printf("write Data: %v \n", randNum)
	}
	close(intChan)
}

// 读线程
func readData() {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(v, "~~~~~~~~~~")

	}
	// 读取完数据，将标志位放到管道中
	exitChan <- true
	close(exitChan)
}
