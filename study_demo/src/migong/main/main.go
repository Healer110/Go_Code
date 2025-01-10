package main

import "fmt"

func showMap(m *[8][7]int) {
	for r := 0; r < 8; r++ {
		for c := 0; c < 7; c++ {
			fmt.Printf("%d ", m[r][c])
		}
		fmt.Println()
	}
}

// 编写一个函数，找到出路
// 参数含义：地图，对地图上的哪个点就行测试, 即起始点
func SetWay(myMap *[8][7]int, r int, c int) bool {
	// 如果map[6][5] = 2就找到出路了
	if myMap[6][5] == 2 {
		fmt.Println("找到出路了。。。。")
		return true
	} else {
		if myMap[r][c] == 0 {
			// 假设这个点是通的，但是需要探测才可以知道结果，策略是下右上左
			myMap[r][c] = 2
			if SetWay(myMap, r+1, c) { // 向下
				return true
			} else if SetWay(myMap, r, c+1) { // 向右
				return true
			} else if SetWay(myMap, r-1, c) { // 向上
				return true
			} else if SetWay(myMap, r, c-1) { // 向左
				return true
			} else { // 死路一条，假设是错误的，将这个点设置为3
				myMap[r][c] = 3
				return false
			}

		} else { // 说明这个点是墙，不能走
			return false
		}
	}
}

func main() {
	// 创建一个二维数组，模拟迷宫
	// 规则
	// 如果元素的值为1表示墙
	// 如果元素的值为0表示还未探测过的路径
	// 如果元素的值为2表示一个通路
	// 如果元素的值为3表示走过的路，但是走不通
	var map1 [8][7]int

	// 先把地图的最上，最下标记为墙
	for i := 0; i < 7; i++ {
		map1[0][i] = 1
		map1[7][i] = 1
	}

	// 最左边，最右边标记为1
	for i := 0; i < 8; i++ {
		map1[i][0] = 1
		map1[i][6] = 1
	}

	map1[3][1], map1[3][2] = 1, 1

	SetWay(&map1, 1, 1)

	showMap(&map1)
}
