package main

import "fmt"

func main() {
	// var fbArr []int = make([]int, 10)
	// for i := 1; i < 11; i++ {
	// 	fbArr[i-1] = fbn(i)
	// }

	// for _, val := range fbArr {
	// 	fmt.Println("value =", val)
	// }

	// fbn2(10)

	var arr = []int{100, 66, 200, 2, 69}
	fmt.Println("before sorting:", arr)
	bubbleSort(arr)
	fmt.Println("after sorting:", arr)

}

// 编写一个函数 fbn(n int)，要求完成
// 可以接收一个n int
// 能够将斐波那契的数列放到切片中

// func fbn(n int) int {
// 	if n == 1 || n == 2 {
// 		return 1
// 	} else {
// 		return fbn(n-1) + fbn(n-2)
// 	}

// }

func fbn2(n int) {
	var fbSlice = make([]int, 10)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			fbSlice[i] = 1
		} else {
			fbSlice[i] = fbSlice[i-2] + fbSlice[i-1]
		}
	}
	fmt.Println("fbSlice =", fbSlice)
}

// 冒泡排序，从小到大排序
func bubbleSort(arr []int) {
	arrLength := len(arr)
	var tmp int
	for i := 0; i < arrLength-1; i++ {
		for j := 0; j < (arrLength - 1 - i); j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}
