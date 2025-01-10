package main

import "fmt"

func main() {
	gotoDemo()
	returnDemo()
}

func gotoDemo() {
	// 演示goto语句的使用
	fmt.Println("ok1")
	fmt.Println("ok2")
	goto lable1
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
lable1:
	fmt.Println("ok6")
	fmt.Println("ok7")
	fmt.Println("ok8")
	fmt.Println("ok9")
}

func returnDemo() {
	for i := 0; i < 10; i++ {
		if i == 8 {
			return
		}
		fmt.Println("i =", i)
	}
}
