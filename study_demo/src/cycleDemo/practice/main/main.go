package main

import (
	"fmt"
)

func main() {
	practice01()
	practice02()
}

/*
打印1~100之间所有是9的倍数的整数的个数以及总和
*/
func practice01() {
	var totalNumber int = 0
	var sumNumber int = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			totalNumber++
			sumNumber += i
		}
	}
	fmt.Println("9的倍数的数字的个数：", totalNumber)
	fmt.Println("9的倍数的数字的总和：", sumNumber)
}

/*
完成下面的表达式输出
0 + 6 = 6
1 + 5 = 6
2 + 4 = 6
3 + 3 = 6
5 + 1 = 6
6 + 0 = 6
*/
func practice02() {
	var sumNumber int = 6
	for i := 0; i <= sumNumber; i++ {
		fmt.Printf("%d + %d = %d \n", i, sumNumber-i, sumNumber)

	}
}
