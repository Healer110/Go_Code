package main

import "fmt"

func main() {
	res := practice01(7)
	fmt.Println("res =", res)

	sum := multiParas(1, 10, 100)
	fmt.Println("sum =", sum)

	fmt.Println("sum =", mySum(1, 2))

	var n1 int = 10
	var n2 int = 20
	swapNumber(&n1, &n2)
	fmt.Printf("n1 = %d, n2 = %d\n", n1, n2)
}

// 斐波那契
func practice01(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return practice01(n-1) + practice01(n-2)
	}
}

// 函数支持多个参数
func multiParas(args ...int) (sum int) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func mySum(n1, n2 float32) float32 {
	fmt.Printf("n1 type = %T\n", n1)
	return n1 + n2
}

// 互换两个数的值，改变之前的变量，使用指针实现
func swapNumber(n1, n2 *int) {
	var tmp int
	tmp = *n1
	*n1 = *n2
	*n2 = tmp
}
