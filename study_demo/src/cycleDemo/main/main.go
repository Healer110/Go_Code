package main

import (
	"fmt"
)

func main() {
	forDemo1()
	forDemo2()
	forDemo3()

	forDemo4()
	for_range_demo()

	whileDemo()
	doWhileDemo()
}

func forDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Println("你好，Golang...", i)
	}
}

// for循环的第二种写法
func forDemo2() {
	j := 10
	for j < 20 {
		fmt.Println("for 循环的第二种写法：", j)
		j++
	}

}

// for循环的第三种写法-->死循环的写法，等价于 for ;; {}
func forDemo3() {
	k := 1
	for { // 等价于 for ; ; {}
		if k <= 10 {
			fmt.Println("死循环...", k)
			k++
			continue
		}
		fmt.Println("执行完毕，break退出...")
		break
	}
}

// 遍历字符串，传统方式，传统方式是按照每个字节取的，带中文的需要使用切片
func forDemo4() {
	var str string = "Hello, world北京"
	fmt.Println("for 循环, 按照传统的方式输出字符串")
	for i := 0; i < len(str); i++ {
		fmt.Printf("字符串-->字符：%c\n", str[i])
	}

	fmt.Println("for 循环, 按照传统的方式输出带中文的字符串")
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		fmt.Printf("字符串-->字符：%c\n", str1[i])
	}
}

// for-range方式遍历字符串
func for_range_demo() {
	var str string = "Hello, world北京"
	fmt.Println("for-range方式输出字符串")
	for idx, val := range str {
		fmt.Printf("字符串-->字符：index=%d, value=%c\n", idx, val)
	}
}

// Go语言中，没有while， do-while循环，但是可以使用for循环的死循环替代
func whileDemo() {
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Println("Hello, world... i =", i)
		i++
	}
}

func doWhileDemo() {
	fmt.Println("do-while Demo......")
	i := 0
	for {
		fmt.Println("Hello, world... i =", i)
		i++
		if i == 10 {
			break
		}
	}
}
