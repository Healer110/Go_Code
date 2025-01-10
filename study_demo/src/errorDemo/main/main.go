package main

import (
	"errors"
	"fmt"
)

func main() {
	// errorDemo01()
	// fmt.Println("遇到错误后，继续执行的代码。。。。。")

	testErr()
	fmt.Println("遇到错误后，继续执行的代码。。。。。")
}

func errorDemo01() {

	// 使用defer 跟 recover处理错误
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error =", err)
			// 捕获到错误后，可以在治理进一步处理
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res =", res)

}

// 自定义错误类型
func selfErrorDefine(name string) (err error) {
	if name == "config.ini" {
		// 读取配置文件
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

func testErr() {
	err := selfErrorDefine("config1.ini")
	if err != nil {
		// 读取文件发生错误，就输出错误，并终止程序
		panic(err)
	}

	fmt.Println("testErr 后面的代码继续执行...")
}
