package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// practice01()
	findMax()
}

// 创建一个byte类型的26个元素的数组，分别放置A-Z，使用for循环访问所有元素并打印出来
func practice01() {
	var arr [26]byte
	arr[0] = 'A'
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[0] + byte(i)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%c \t", arr[i])
		if (i+1)%5 == 0 {
			fmt.Println()
		}
	}
}

// 求出数组中的最大值
func findMax() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	for x := len(arr) - 1; x >= 0; x-- {
		fmt.Printf("%d ", arr[x])
	}

}
