package main

import "fmt"

func main() {
	var s Stu
	s.Name = "Jack"
	// s对象调用的方法的时候，会自动将自身传递进去
	s.demo01()
	s.speak()
	res := s.jisuan(100)
	fmt.Println("res =", res)
}

type Stu struct {
	Name string
	Age  int8
}

// 针对指定数据类型的叫做方法, 给Stu结构体绑定一个方法
func (s Stu) demo01() {
	fmt.Println(s.Name)
}

// 为Stu结构体绑定的另外一个方法
func (s Stu) speak() {
	fmt.Printf("%v正在说话...\n", s.Name)
}

func (s Stu) jisuan(n int) int {
	var tmp int = 0
	for i := 1; i <= n; i++ {
		tmp += i
	}
	fmt.Println("total number =", tmp)
	return tmp
}
