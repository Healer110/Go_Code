package main

import "fmt"

// 在Go的初始化时，先初始化全局变量，在执行init函数，再执行main函数
var a = test()

func test() int {
	fmt.Println("test function...")
	return 90
}

// 通常可以在init函数中，完成一个初始化的工作，运行时，会首先调用init函数
func init() {
	fmt.Println("init function...")
	fmt.Println("a =", a)
}

func main() {
	fmt.Println("main function...")
}
