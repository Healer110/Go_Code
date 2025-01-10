package main

import (
	"factoryDemo/model"
	"fmt"
)

func main() {
	// var stu = model.student{
	// 	Name:  "tom",
	// 	Score: 56.5,
	// }

	stu := model.NewStudengt("Tom", 98.7)

	fmt.Println(*stu)
}
