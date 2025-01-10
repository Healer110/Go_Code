package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "指定用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "", "主机，默认为空")
	flag.IntVar(&port, "p", 3306, "端口号，默认3306")

	// 解析指令，必须转换，不然会失败
	flag.Parse()

	fmt.Printf("user = %v, pwd = %v, host = %v, port = %v \n", user, pwd, host, port)

}
