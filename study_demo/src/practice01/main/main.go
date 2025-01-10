package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int32 = 23
	var num2 int32 = 40
	demo1(num1, num2)

	var num3, num4 float32
	num3 = 99.9
	num4 = 19.5
	demo2(num3, num4)

	var num5 int32 = 14
	var num6 int32 = 15
	demo3(num5, num6)

	demo4(12.2, 14.5, 1.0)

	demo5()

}

func demo1(num1 int32, num2 int32) {
	temp := num1 + num2
	if temp >= 50 {
		fmt.Println("The sum of two number is greater than 50...")
	}
}

func demo2(num1 float32, num2 float32) {
	if num1 > 10.0 && num2 < 20.0 {
		fmt.Println("The sum of two number is", num1+num2)
	} else {
		fmt.Println("Error....")
	}
}

func demo3(num1 int32, num2 int32) {
	temp := num1 + num2
	if temp%3 == 0 && temp%5 == 0 {
		fmt.Println("可以被3跟5整除")
	} else {
		fmt.Println("不可以被3跟5整除")
	}
}

func demo4(a float64, b float64, c float64) {
	flag := math.Pow(b, 2) - 4*a*c
	if flag > 0 {
		x1 := (-b + math.Sqrt(flag)) / 2 * a
		x2 := (-b - math.Sqrt(flag)) / 2 * a
		fmt.Println("两个解:", x1, x2)
	} else if flag == 0 {
		x1 := (-b + math.Sqrt(flag)) / 2 * a
		fmt.Println("一个解:", x1)
	} else {
		fmt.Println("无解")
	}
}

func demo5() {
	var achievement float32
	var gender string = "male"
	fmt.Println("请输入比赛队员的成绩：")
	fmt.Scanln(&achievement)

	if achievement < 8.0 {
		if gender == "male" {
			fmt.Println("恭喜你，成功进入男子组决赛。")
		} else if gender == "female" {
			fmt.Println("恭喜你，成功进入女子组决赛。")
		}
	} else {
		fmt.Println("成绩不达标，未进入决赛。")
	}
}
