package main

import "fmt"

// 演示闭包
func AddUpper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(i int) int {
		n = n + i
		str += string(36)
		fmt.Println("str =", str)
		return n
	}
}

func main() {
	f := AddUpper()
	res := f(1)
	fmt.Println("res =", res)
	res = f(2)
	fmt.Println("res =", res)
}
