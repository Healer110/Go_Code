package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// practice01()
	practice02()
}

// 猜数字游戏，生成1-100之内的随机数，用户有十次机会，猜中提示猜中了，机会使用完，未猜中提示未猜中游戏结束
func practice01() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	generateNUmber := rand.Intn(100) + 1
	fmt.Println("随机生成的数字为:", generateNUmber)
	var inputNumber int
	for i := 0; i < 10; i++ {
		fmt.Println("请猜测随机生成的数字：")
		fmt.Scanln(&inputNumber)
		if inputNumber == generateNUmber {
			fmt.Println("恭喜你，猜对了...")
			return
		} else if inputNumber > generateNUmber {
			fmt.Println("猜大了...")
			continue
		} else {
			fmt.Println("猜小了...")
			continue
		}
	}
	fmt.Println("十次机会已用完，未猜对...")

}

// 编写一个函数，输出100以为的所有素数(), 每行显示5个，并求和
// 素数：只能被1跟自己整除
func practice02() {
	var sum int
	var str string = ""
lable1:
	for i := 1; i <= 100; i++ {
		if i == 1 || i == 2 {
			sum += i
			str += strconv.Itoa(i) + " "
		} else {
			for j := 2; j < i; j++ {
				if i%j == 0 {
					continue lable1
				}
			}

			sum += i
			str += strconv.Itoa(i) + " "
		}
	}

	// 遍历完成打印
	strArr := strings.Split(strings.TrimSpace(str), " ")
	fmt.Println("1~100的素数如下：")
	for i := 1; i < len(strArr)+1; i++ {
		tmp, err := strconv.Atoi(strArr[i-1])
		if err != nil {
			fmt.Println("字符串中有错误...")
			return
		}
		fmt.Printf("%v\t", tmp)
		if i%5 == 0 {
			fmt.Println()
		}
	}
}
