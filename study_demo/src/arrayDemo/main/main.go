package main

import "fmt"

func main() {
	test01()
}

func test01() {
	// 定义一个数组
	var hens [6]float64
	hens[0] = 5.0
	hens[1] = 5.0
	hens[2] = 5.0
	hens[3] = 5.0
	hens[4] = 5.0
	hens[5] = 5.0

	for _, v := range hens {
		fmt.Println("Value =", v)
	}

}
