package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 4}
	a = point
	var b Point
	// 下面的赋值会报错，需要使用类型断言
	// b = a
	// 使用下面的断言：表示判断a是否指向Point类型的变量，如果是就转成Point类型
	// 并赋值给b变量，否则报错
	b = a.(Point)
	fmt.Println(b)

	fmt.Println("take out function...")
	judgeBeforeAssert()

	var n1 int = 11
	var n2 float64 = 1.111
	var n3 string = "hsagasg"
	var n4 bool = true
	judgeInputType(n1, n2, n3, n4)
}

func judgeBeforeAssert() {
	var x interface{}
	var f float32 = 1.345
	x = f
	// y, ok := x.(float64)
	// 带检测的断言
	// y, ok := x.(float32)
	// if ok {
	// 	fmt.Println("convert success...")
	// 	fmt.Println(y)
	// } else {
	// 	fmt.Println("convert fail...")
	// }

	if y, ok := x.(float32); ok {
		fmt.Println("convert success...")
		fmt.Println(y)
	} else {
		fmt.Println("convert fail...")
	}

}

func judgeInputType(paras ...interface{}) {
	for _, v := range paras {
		switch v.(type) {
		case bool:
			fmt.Printf("value=%v, type=bool \n", v)
		case int, int64:
			fmt.Printf("value=%v, type=int \n", v)
		case float32:
			fmt.Printf("value=%v, type=float32 \n", v)
		case float64:
			fmt.Printf("value=%v, type=float64 \n", v)
		case string:
			fmt.Printf("value=%v, type=string \n", v)
		default:
			fmt.Printf("类型不确定。。。")
		}

	}
}
