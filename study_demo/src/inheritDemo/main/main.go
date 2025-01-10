package main

import "fmt"

/*
继承：提取不同对象实例的相同部分，组成最顶层的父类
每个实例继承父类的属性，并添加自己独有的部分
go实现的方法是嵌入一个匿名结构体
匿名字段也可以是基本数据类型
*/
type Student struct {
	name  string
	score int
}

func (stu *Student) ShowInfo() {
	fmt.Println("Student name is", stu.name)
	fmt.Println("score is", stu.score)
}

func (stu *Student) SetScore(score int) {
	stu.score = score
}

// 小学生结构体，继承student
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

// 匿名字段是基本数据类型
type A struct {
	Student
	int
}

func main() {
	p := &Pupil{}
	p.Student.name = "Tom"
	p.Student.score = 100
	p.Student.ShowInfo()
	p.Student.SetScore(200)
	p.Student.ShowInfo()
	p.testing()

	// 匿名字段是基本数据类型，定义以及使用
	var a A
	a.name = "March"
	a.score = 99
	a.int = 204
	fmt.Println("匿名字段是基本数据类型..")
	fmt.Println(a)

}
