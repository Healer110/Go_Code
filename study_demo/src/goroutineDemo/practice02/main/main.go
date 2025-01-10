package main

import (
	"fmt"
)

/*
启动一个协程，将1-2000的数放入到一个channel中，比如numChan
启动8个协程，从numChan取出数，并计算1+....n的值，并存放到resChan
最后8个协程完成工作后，再遍历resChan，显示结果res[1]=1, ....res[10]=55
注意，考虑resChan chan int是否合适？
*/

// 定义一个channel，用来存放每个线程完成后存在一个true用来表示结束
var flagChan chan bool = make(chan bool, 8)

func writeNumber(nChan chan int) {
	for i := 1; i <= 2000; i++ {
		nChan <- i
	}
	close(nChan)
}

// 读取数据
func readNumber(nChan chan int, resCh chan int) {
	for {
		res := 0
		v, ok := <-nChan
		if !ok {
			break
		}

		for i := 1; i <= v; i++ {
			res += i
		}
		resCh <- res
	}
	flagChan <- true
}

func main() {
	var numChan chan int = make(chan int, 2000)
	var resChan chan int = make(chan int, 2000)

	go writeNumber(numChan)
	for i := 0; i < 8; i++ {
		go readNumber(numChan, resChan)
	}

	for i := 0; i < 8; i++ {
		<-flagChan
	}
	close(flagChan)
	close(resChan)
	for val := range resChan {
		fmt.Println(val)
	}

	// for i := 0; i < 20; i++ {
	// 	num := <-resChan
	// 	fmt.Printf("res[%v] = %v \n", i, num)
	// }

	fmt.Println("run finish...")
}
