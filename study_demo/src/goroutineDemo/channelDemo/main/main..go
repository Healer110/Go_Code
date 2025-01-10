package main

import "fmt"

// 管道也可以存放任意数据类型，只需要将指定的数据类型定义为空接口即可
func main() {
	// 创建管道：管道类似于队列，FIFO，管道本身具备线程安全特性(这个是由GO底层维护，无需手动操作)
	var intChan chan int = make(chan int, 3)
	// 管道是引用类型
	fmt.Println(intChan)
	fmt.Printf("channel 变量本身的地址=%p\n", &intChan)

	//向管道写入数据
	intChan <- 10
	num := 222
	intChan <- num
	fmt.Println(len(intChan), cap(intChan))

	// 从管道中取数据
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)
	fmt.Println(len(intChan), cap(intChan))

	// 在灭有使用协程的情况下，如果channel中没有数据，再去获取就会报错
	num3 := <-intChan
	num4 := <-intChan
	num5 := <-intChan
	fmt.Println(num3, num4, num5)
}
