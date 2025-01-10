package main

import (
	"fmt"
)

func main() {
	// switch_demo()

	// switch_demo2()

	switch_demo3()
}

func switch_demo() {
	var input_char byte
	fmt.Println("请输入一个字符：a,b,c,d,e")
	fmt.Scanf("%c", &input_char)

	switch input_char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}

}

func switch_demo2() {
	var score float64 = 80

	switch {
	case score > 80 && score <= 100:
		fmt.Println("合格")
	case score <= 80:
		fmt.Println("不合格")
	}
}

func switch_demo3() {
	var month int8
	fmt.Println("请出入月份：")
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	}
}
