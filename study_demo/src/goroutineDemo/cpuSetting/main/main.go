package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 获取cpu的数量
	// num := runtime.NumCPU()
	// fmt.Println("CPU 的数量:", num)

	// // 设置使用CPU的数量,在1.8版本以后，程序默认运行在多核上，可以不用手动设置了
	// runtime.GOMAXPROCS(4)

	for i := 1; i <= 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 5)

	// 输出结果
	lock.Lock()
	for key, v := range numberMap {
		fmt.Printf("%v: %v \n", key, v)
	}
	lock.Unlock()

}

// 编写一个函数计算各个数的阶乘，并存放到map中
// 启动多个协程访问map，将结果保存
// map是个全局变量

var numberMap map[int]int = make(map[int]int)
var lock sync.Mutex

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	numberMap[n] = res
	lock.Unlock()
}
