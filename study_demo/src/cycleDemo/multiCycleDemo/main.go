package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// demo01()
	randomDemo()
}

/*
统计三个班成绩情况，每个班有5名同学，求出每个班的平均分和所有班级的平均分
学生的成绩从键盘输入
*/
func demo01() {
	var classNum = 3
	var students = 5
	var sumScore float64 = 0.0
	var inputNumber float64 = 0.0
	var passStudentNum int = 0

	for i := 1; i <= classNum; i++ {
		var classAverageScore float64 = 0.0
		for j := 1; j <= students; j++ {
			fmt.Printf("请输入%d班, 学生%d的考试分数：\n", i, j)
			fmt.Scanln(&inputNumber)
			classAverageScore += inputNumber
			if inputNumber > 60 {
				passStudentNum++
			}
		}
		sumScore += classAverageScore
		fmt.Printf("%d班的平均分是：%f\n", i, classAverageScore/float64(students))
	}
	fmt.Printf("%d班的平均成绩是：%f\n", classNum, sumScore/(float64(classNum)*float64(students)))
	fmt.Printf("%d班的总及格人数：%d\n", classNum, passStudentNum)
}

// 随机生成1-100的一个数，直到生成了99这个数，看看你一共用了几次
func randomDemo() {
	var num uint64 = 0
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		res := rand.Intn(100) + 1
		num++
		if res == 99 {
			break
		}
	}
	fmt.Println("找到随机数99，使用的次数:", num)
}
