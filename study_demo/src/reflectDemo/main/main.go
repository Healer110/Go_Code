package main

import (
	"fmt"
	"reflect"
)

// 使用反射获取变量的类型以及值
func reflectTest01(b interface{}) {
	// 获取类型reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)

	// 获取值reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v, rValue type=%T \n", rValue, rValue)

	// 转为int
	n2 := 2 + rValue.Int()
	fmt.Println("n2 =", n2)

	// 将反射的返回值转为空接口类型
	iValue := rValue.Interface()
	// 通过断言将空接口类型转为int类型
	intNUm := iValue.(int)
	fmt.Println("intNUm =", intNUm)
}

// 对结构体进行反射操作
func reflectTest02(s interface{}) {
	// 获取类型reflect.Type
	rType := reflect.TypeOf(s)
	fmt.Println("rType =", rType)

	// 获取值reflect.Value
	rValue := reflect.ValueOf(s)
	fmt.Printf("rValue=%v, rValue type=%T \n", rValue, rValue)

	// 转为空接口
	iValue := rValue.Interface()
	fmt.Printf("rValue=%v, rValue type=%T \n", iValue, iValue)
	// 断言转为真实的结构体类型
	stu := iValue.(Student)
	fmt.Println(stu.Name, stu.Age)

}

type Student struct {
	Name string
	Age  uint8
}

func main() {
	// var n int = 100
	// reflectTest01(n)

	// 结构体反射操作
	s := Student{
		Name: "John",
		Age:  33,
	}
	reflectTest02(s)

}
