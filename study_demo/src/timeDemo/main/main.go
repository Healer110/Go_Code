package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// timeDemo01()
	// timeDemo02()

	// 函数执行时间统计
	// 先获取时间戳
	start := time.Now().Unix()
	timeCalc()
	end := time.Now().Unix()
	fmt.Printf("函数的运行时间: %v秒\n", end-start)
}

// 时间相关的函数
func timeDemo01() {
	var t time.Time = time.Now()
	fmt.Println("now time =", t)

	// 格式化时间
	fmt.Printf("年=%v\n", t.Year())
	// fmt.Printf("月=%v\n", t.Month())
	fmt.Printf("月=%v\n", int(t.Month()))
	fmt.Printf("日=%v\n", t.Day())
	fmt.Printf("时=%v\n", t.Hour())
	fmt.Printf("分=%v\n", t.Minute())
	fmt.Printf("秒=%v\n", t.Second())

	// 格式化输出
	// 字符串中的值是固定的，不能变，但是可以按照指定的间隔符，以及单独指定每个字段
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("2006/01/02 15:04:05"))
	fmt.Println(t.Format("2006/01/02"))

}

func timeDemo02() {
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
		time.Sleep(time.Second * 3)
	}
}

// 统计函数执行的时间
func timeCalc() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
