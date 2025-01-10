package main

import "fmt"

// 放入数据的线程
func putNUm(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 去除数据，计算素数的线程
func primeNumber(intchan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intchan
		if !ok { // 取不到了
			break
		}
		// 判断素数
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 { // 不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("协程获取不到数据了，退出...")
	exitChan <- true

}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 存放计算出来的素数
	exitChan := make(chan bool, 4)    // 标识4个线程的退出

	// 开启协程放入数据
	go putNUm(intChan)

	// 开启4个线程计算素数
	for i := 0; i < 4; i++ {
		go primeNumber(intChan, primeChan, exitChan)
	}

	// 获取4个协程退出的标志
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// close(exitChan)
		close(primeChan)
	}()

	// 遍历素数channel
	// 1-8000的素数：
	fmt.Println("1-8000素数的个数：", len(primeChan))
	fmt.Println("1-8000的素数：")
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
	// for v := range primeChan {
	// 	fmt.Println(v)
	// }

}
