package main

import "fmt"

// 使用全局变量，定义一个匿名函数
var (
	func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 使用匿名函数的方式，求两个数的和
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res =", res)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res = a(20, 10)
	fmt.Println("res =", res)

	res = func1(10, 20)
	fmt.Println("res =", res)

}
