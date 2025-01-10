package main

import "fmt"

func main() {
lable1: // for循环设置一个标签，break可以跳出指定的标签标识的for循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break lable1
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

}
