package main

import "fmt"

func main() {
	demo01()
}

func demo01() {
	// 二位数组的声明
	// 下面的声明定义一个二维数组，有4个一维数组， 每个一维数组有5个int类型的元素
	var arr [4][5]int
	fmt.Println(arr)
	arr[1][2] = 100
	for _, val := range arr {
		fmt.Println(val)
	}

}
