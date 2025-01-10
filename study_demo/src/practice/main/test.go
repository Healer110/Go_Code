package main

import (
	"fmt"
	"practice/model"
	"strconv"
	"unsafe"
)

// 定义全局变量
var GNUM1 = 100
var GNUM2 = 200
var GNAME1 = "Tom"

// 一次性声明
var (
	GNUM3  = 300
	GNUM4  = 400
	GNAME2 = "Tom Tick"
)

func main() {
	var num1 int
	fmt.Println(num1)

	// := 定义并赋值
	/*
		var name string
		name = "dgsgasg"
	*/
	name := "dgsgasg"
	fmt.Println("name= ", name)

	// 自动推导变量类型
	var floatNum = 12.12
	fmt.Println("floatNum = ", floatNum)

	// 一次性多变量声明Demo
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	// 一次性声明不同类型的变量并赋值
	// 推导的方式 name1, age, score := "Lili", 23, 98.5
	var name1, age, score = "Lili", 23, 98.5
	fmt.Println(name1, age, score)

	// 打印全局变量
	fmt.Println(GNUM1, GNUM2, GNAME1)
	fmt.Println(GNUM2, GNUM4, GNAME2)

	// 整数类型
	var kk uint8 = 255
	fmt.Println("kk=", kk)
	fmt.Printf("类型 %T\n", kk)
	fmt.Printf("占用空间 %d\n", unsafe.Sizeof(kk))

	var ab = "ab"
	fmt.Println("char ab =", ab)

	// 基本数据类型跟string类型之间的转换
	var bNum1 int64 = 9999
	var bNum2 float64 = 222.555
	var bFlag = true
	var bChar = 'a'
	var str1 string

	str1 = fmt.Sprintf("%d", bNum1)
	fmt.Printf("str1=%v, type=%T\n", str1, str1)

	str1 = fmt.Sprintf("%f", bNum2)
	fmt.Printf("str1=%v, type=%T\n", str1, str1)

	str1 = fmt.Sprintf("%t", bFlag)
	fmt.Printf("str1=%v, type=%T\n", str1, str1)

	str1 = fmt.Sprintf("%c", bChar)
	fmt.Printf("str1=%q, type=%T\n", str1, str1)

	var str2 string = "false"
	var b bool
	b, _ = strconv.ParseBool(str2)
	fmt.Printf("after strconv b = %v, b's type = %T\n", b, b)

	var str3 string = "345634.34636"
	convNum_floag, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("after strconv convNum_floag = %v, convNum_floag's type = %T\n", convNum_floag, convNum_floag)

	// 复杂数据类型：指针
	var number1 int = 666
	var ptr *int = &number1
	fmt.Printf("number addr=%v, ptr addr=%v, ptr point addr=%v\n", &number1, &ptr, ptr)
	*ptr = 999
	fmt.Printf("ptr value=%v\n", *ptr)
	fmt.Printf("ptr value=%v\n", number1)

	//导入自定义的包
	fmt.Println(model.ModelName)

	// 不使用中间变量的两个值之间的互换操作
	res1, res2 := calc01()
	fmt.Printf("a=%v, b=%v \n", res1, res2)

	fmt.Println(-2 ^ 3)

}

func calc01() (int, int) {
	var a int = 10
	var b int = 20
	a += b
	b = a - b
	a = a - b
	return a, b
}
