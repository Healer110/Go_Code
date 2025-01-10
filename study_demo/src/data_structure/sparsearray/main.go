package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ValueNode struct {
	row int
	col int
	val int
}

func main() {
	// 创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 1表示黑子
	chessMap[2][3] = 2 // 2表示白子

	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

	// 转为稀疏数组
	// 遍历二维数组，如果发现数据不为零，就创建一个node结构体，并将其放入切片中
	var sparseArr []ValueNode

	// 标准的稀疏数组，含有表示记录原始的二维数组的大小
	valNode := ValueNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建节点
				valNode = ValueNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	// 恢复数据, 从chessMap.data文件中恢复
	var chessMap2 [][]int

	filePath := "src/data_structure/sparsearray/chessMap.data"
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0664)
	if err != nil {
		fmt.Println("打开文件异常...")
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	firstLineFlag := true
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.Trim(str, "\r\n")
		numList := strings.Split(str, " ")
		fmt.Println(numList)
		row, _ := strconv.Atoi(numList[0])
		col, _ := strconv.Atoi(numList[1])
		val, _ := strconv.Atoi(numList[2])
		fmt.Println(row, col, val)
		if firstLineFlag {
			chessMap2 = make([][]int, row)
			for i := range chessMap2 {
				chessMap2[i] = make([]int, col)
			}
			for i := 0; i < row; i++ {
				for j := 0; j < col; j++ {
					chessMap2[i][j] = val
				}
			}
			firstLineFlag = false
			continue
		}

		chessMap2[row][col] = val

		// fmt.Print(str)
	}

	for _, v := range chessMap2 {
		for _, val := range v {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

}
