package main

import "fmt"

func main() {
	demo01()
	demo02()
}

/*
100以内的数求和，求出当和第一次大于20的当前数
*/
func demo01() {
	var sum int16 = 0
	for x := 0; x <= 100; x++ {
		sum += int16(x)
		if sum > 20 {
			fmt.Printf("sum=%d, first number=%d\n", sum, x)
			break
		}
	}
}

/*
	实现登录验证，有三次机会，如果用户名为“张无忌”，密码“888”提示登录成功，否则提示还有几次机会
*/
func demo02() {
	var username string
	var password string
	for x := 1; x < 4; x++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&username)
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		if username == "张无忌" && password == "888" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Println("验证失败，还可以输入的次数：", 3-x)
		}

	}
}
