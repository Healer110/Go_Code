package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
开一个协程writeDataToFile,随机成成1000个数，存放到文件中
当写协程完成后，让sort协程从文件中读取1000个数，并将排序后的数据写入到另外一个文件
考察点：协程管道+文件的综合使用
功能扩展：
开10个协程writeDataToFile，每个协程随机生成1000个数，存放到10个文件中
当10个文件都生成后，让10个sort协程，读取10个文件中的数据，排序完成后，写入10个文件中
*/

// 写文件
func writeDataToFile(filePath string, exitChan chan bool) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0664)
	defer file.Close()
	writer := bufio.NewWriter(file)

	for i := 0; i < 1000; i++ {
		data := rand.Intn(1000) + 1
		_, err := writer.WriteString(strconv.Itoa(data) + "\n")
		if err != nil {
			fmt.Println("write data error: ", err)
			return
		}
	}
	writer.Flush()
	exitChan <- true
	close(exitChan)
}

// 读协程
func readDataFromFile(filePath string, intSlice []int, sortFilePath string, eixtChan chan bool) {
	file, _ := os.OpenFile(filePath, os.O_RDONLY, 0664)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		str = strings.TrimRight(str, "\n")
		if err == io.EOF {
			break
		}
		number, _ := strconv.Atoi(str)
		intSlice = append(intSlice, number)
	}
	sort.Ints(intSlice)

	wFile, _ := os.OpenFile(sortFilePath, os.O_APPEND|os.O_CREATE, 0664)
	defer wFile.Close()
	writer := bufio.NewWriter(wFile)
	for _, v := range intSlice {
		writer.WriteString(strconv.Itoa(v) + "\n")
	}
	writer.Flush()
	eixtChan <- true
	close(eixtChan)

}

func main() {
	var exitChan chan bool = make(chan bool, 1)
	var sortExitChan chan bool = make(chan bool, 1)
	randomFile := "random.txt"
	sortFile := "sort.txt"
	var sortSlice []int
	go writeDataToFile(randomFile, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	go readDataFromFile(randomFile, sortSlice, sortFile, sortExitChan)
	for {
		_, ok := <-sortExitChan
		if !ok {
			break
		}
	}
	fmt.Println("finish...")
}
