package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// demo01()
	demo02()
}

// 定义结构体
type Person struct {
	name    string
	age     int
	address string
}

// 使用结构体
func demo01() {
	// 结构体类似声明一个对象
	var p1 Person
	p1.name = "zzz"
	p1.age = 23
	p1.address = "北京"

	var p2 Person = Person{"aaa", 33, "上海"}
	var p3 *Person = new(Person)
	(*p3).name = "bbb"
	(*p3).age = 28
	// 结构体指针在底层做了优化，赋值变量或者更改变量值的时候，可以不使用*，底层代码会自动添加
	p3.address = "深圳"

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(*p3)
}

// 定义字段后面加入结构体标签，这样在序列化时，会使用标签替换为字段名（使用的技术是反射）
type Stu struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

// struct 每个字段上可以写上一个tag, 该tag可以通过反射机制获取，常见的使用场景是序列化和反序列化
func demo02() {
	s1 := Stu{"zzzz", 33}
	jsonStr, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("error...")
	} else {
		fmt.Println("jsonStr =", string(jsonStr))
	}
}
