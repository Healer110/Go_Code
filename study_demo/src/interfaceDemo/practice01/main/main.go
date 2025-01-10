package main

import "fmt"

/*
接口中包含接口
*/

type A interface {
	testA()
}

type B interface {
	testB()
}

type C interface {
	A
	B
	testC()
}

type Stu struct {
}

func (student Stu) testA() {
	fmt.Println("testA()....")
}

func (student Stu) testB() {
	fmt.Println("testB()....")
}

func (student Stu) testC() {
	fmt.Println("testC()....")
}

// func (s *Stu) testC() {
// 	fmt.Println("testC()....")
// }

func main() {
	var stu Stu
	var c C = stu
	c.testA()
	c.testB()
	c.testC()
}
