package main

import (
	"fmt"
	"math"
)

func main() {
	var c Circle = Circle{2}
	res := (&c).area()
	// res := c.area()
	fmt.Printf("area = %.2f\n", res)
	// 这里传入时，要传入指针变量，因为Stirng()方法绑定的是该类型的指针
	fmt.Println(&c)

	// ===================
	meth := MethodUtils{"Method_struct"}
	meth.methodPrint(6, 8)
}

/*
声明一个结构体Circle，字段为radium
声明一个方法area和Circle绑定可以返回面积
*/
type Circle struct {
	radium float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.radium, 2)
	// return math.Pi * math.Pow((*c).radium, 2)
}

// 如果实现了自定义类型的String()方法，再代用打印函数时，就会自动调用该方法
// 可以应用于日志打印场景，格式化输出字符串
func (c *Circle) String() string {
	str := fmt.Sprintf("radium = [%f]", c.radium)
	return str
}

// 编写结构体（MethodUtils)
// 编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
type MethodUtils struct {
	name string
}

func (meth *MethodUtils) methodPrint(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}


