package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// demo01()
	// demo02()
	// demo03()
	// demo04()
	// demo02()
	// demo05()
	// demo06()
	// demo07()
	// fileIsExist()
	copyFile()
}

// 不带缓冲的demo
func demo01() {
	// 打开一个文件
	file, err := os.Open("D:\\test.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	} else {
		fmt.Println(file)
		err = file.Close()
		if err != nil {
			fmt.Println("关闭文件错误...")
		}
	}
}

// 带缓冲的demo
func demo02() {
	// 读取文件的内容并显示在终端
	// 打开一个文件
	file, err := os.Open("D:\\test.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	}
	// 函数退出时，及时关闭文件句柄，防止内存泄漏
	defer file.Close()

	// 创建一个*Reader，带缓冲的
	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(content)
	}
	fmt.Println("文件读取结束...")
}

// 一次性读取文件内容
func demo03() {
	file := "D:\\test.txt"
	// 没有打开，也没有关闭，ReadFile方法会自动处理
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("read file err =", err)
	}

	fmt.Println(content)
	// 转换为字符串
	fmt.Println(string(content))

}

// 写文件操作示例
func demo04() {
	file := "D:\\test.txt"
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("打开文件异常异常...")
		return
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	res, err := writer.WriteString("new content from program...\n")
	if err != nil {
		fmt.Println("写入异常...")
		return
	}
	// writer写入的时候是先缓存，执行flush后才会落盘
	writer.Flush()
	fmt.Println(res)
	fmt.Println("输出完成...")
}

// 打开一个存在的文件，将原来的内容覆盖成新的内容10句 你好，~~~
func demo05() {
	file := "D:\\test.txt"
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open file error...")
		return
	}
	defer f.Close()
	// 创建缓存writer
	writer := bufio.NewWriter(f)
	for i := 0; i < 10; i++ {
		writer.WriteString("你好，~~~\n")
	}
	// 刷新缓存
	writer.Flush()
}

// 打开一个存在的文件，在原来的内容追加内容 “ABC | ENGLISH"
func demo06() {
	file := "D:\\test.txt"
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开文件异常...")
		return
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	writer.WriteString("ABC | ENGLISH\n")
	writer.Flush()
}

// 打开一个存在的文件，将原来的内容读出来显示在终端，并且追加5句“你好，北京！”
func demo07() {
	file := "D:\\test.txt"
	// 该可读可写模式，写的时候会追加，不会覆盖
	f, err := os.OpenFile(file, os.O_RDWR, 0664)
	if err != nil {
		fmt.Println("打开文件异常...")
		return
	}
	defer f.Close()
	// 1 读取文件内容
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	// 写操作
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString("你好，北京！\n")
	}
	writer.Flush()
}

// 判断文件是否存在
func fileIsExist() {
	file := "D:\\test.txt"
	_, err := os.Stat(file)
	if err == nil {
		fmt.Println("文件存在...")
	} else if os.IsNotExist(err) {
		fmt.Println("文件不存在...")
	} else {
		fmt.Printf("err = %v\n", err)
	}
}

// 将文件从一个目录拷贝到另外一个目录
func copyFile() {
	srcFile := "D:\\test.txt"
	desFile := "D:\\项目备份\\destest.txt"

	f, err := os.OpenFile(desFile, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("打开写文件异常：%v\n", desFile)
		return
	}
	defer f.Close()
	writer := bufio.NewWriter(f)

	rf, err := os.OpenFile(srcFile, os.O_RDONLY, 0664)
	if err != nil {
		fmt.Printf("打开读文件异常：%v\n", srcFile)
		return
	}
	defer rf.Close()
	reader := bufio.NewReader(rf)

	_, err = io.Copy(writer, reader)
	if err != nil {
		fmt.Printf("copy 文件异常：%v\n", err)
	} else {
		fmt.Println("完成copy......")
	}
}
