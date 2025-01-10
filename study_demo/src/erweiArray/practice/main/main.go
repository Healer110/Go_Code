package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	practice01()
}

// 定义一个二维数组，用于保存三个班，每个班5名学生的成绩
// 求出每个班级的平均分以及所有班级的平均分
func practice01() {
	var score [3][5]int8
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			score[i][j] = int8(rand.Intn(100) + 1)
		}
	}

	// 输出每个班学生的成绩
	var classAvgScore float64
	for idx, v := range score {
		fmt.Printf("%d班的学生成绩：", idx+1)
		var avgScore float64 = 0
		for _, s := range v {
			fmt.Printf("%v\t", s)
			avgScore += float64(s)
		}
		classAvgScore += avgScore
		fmt.Printf("%d班的平均分: %v\n", idx+1, avgScore/float64(len(v)))
	}
	fmt.Printf("三个班的平均成绩: %.2f\n", classAvgScore/15)
}
