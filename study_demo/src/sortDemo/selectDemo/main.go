package main

import "fmt"

func main() {
	score := [6]int{10, 34, 19, 100, 80, 789}

	scoreLen := len(score)
	fmt.Println("排序前:", score)
	maxIndex := 0

	for i := 0; i < scoreLen-1; i++ {
		for j := 0; j < scoreLen-1-i; j++ {
			if score[j] < score[j+1] {
				maxIndex = j + 1
			}
		}
		// 记录最大值的index, 最后做一次交换进行
		if maxIndex != scoreLen-i-1 {
			score[maxIndex], score[scoreLen-1-i] = score[scoreLen-1-i], score[maxIndex]
			fmt.Println("交换...", score)
		}
	}
	fmt.Println("排序后:", score)
}
