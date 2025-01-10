package main

import (
	"fmt"
	"strings"
)

func main() {
	// demo01()
	// demo02()
	// demo03()
	// stringSlice()
	modifyString()
}

func demo01() {
	// 切片的基本使用
	var intArr [5]int = [...]int{1, 22, 44, 77, 99}
	sliceArr := intArr[1:3]
	fmt.Printf("intArr type=%T, sliceArr type=%T\n", intArr, sliceArr)
	sliceArr = append(sliceArr, 55, 1000)
	fmt.Println(intArr, sliceArr)
}

func demo02() {
	// 使用make直接定义一个切片
	var sliceA []float64 = make([]float64, 5, 10)
	sliceA[3] = 100
	sliceA[0] = 2
	fmt.Println(sliceA)
	fmt.Println(len(sliceA), cap(sliceA))

	// 方式三创建切片
	var sliceB []string = []string{"Abc", "Tom", "Lisa"}
	// 通过append追加一个元素，要保持类型一致
	sliceB = append(sliceB, "Jack")
	sliceB = append(sliceB, sliceB...)
	fmt.Println(sliceB)
	fmt.Println(len(sliceB), cap(sliceB))
}

func demo03() {
	// 切片拷贝
	var a []int = []int{1, 2, 3, 4, 5}
	var b = make([]int, 6)
	b[0] = 10
	b[1] = 100
	b[5] = 1000
	var slice []int = make([]int, 10)
	fmt.Println(slice)

	copy(slice, a)
	copy(slice[5:], b)
	fmt.Println(slice)
}

func stringSlice() {
	// string底层是一个byte数组，因此string也可以进行切片处理
	var str string = "hello@atguigu"
	strSlice := str[strings.Index(str, "a"):]
	fmt.Println(strSlice)
	fmt.Printf("strSlice type = %T", strSlice)

}

func modifyString() {
	// 字符串修改
	var str string = "hello@atguigu"
	sliceStr := []byte(str)
	sliceStr[0] = 'H'
	// sliceStr[0] = '京'
	fmt.Println(sliceStr)
	str = string(sliceStr)
	fmt.Println(str)
	// 上面的方式只支持英文转换，因为是byte类型，中文按照utf8编码，占用3个字节
	// 如果想重新赋值为中文，使用[]rune切片, []rune是按照字符处理的
	runeStr := []rune(str)
	runeStr[0] = '北'
	str = string(runeStr)
	fmt.Println(runeStr)
	fmt.Printf("%T\n", runeStr)
	fmt.Println(str)
}
