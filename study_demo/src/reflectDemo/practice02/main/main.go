package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 float64
	Num2 float64
}

func (c *Cal) GetSub(name string) {
	fmt.Printf("%s 完成了减法运行，%v - %v = %v \n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func TestReflect(a interface{}) {
	// 当传递进来的是指针类型的时候，获取变量的方法
	// rType := reflect.TypeOf(a)
	// fmt.Println(rType)
	rValue := reflect.ValueOf(a).Elem()
	r := reflect.ValueOf(a)
	fmt.Println(rValue)
	if rValue.Kind() != reflect.Struct {
		fmt.Println("函数需要struct类型...")
		return
	}
	num := rValue.NumField()
	fmt.Println("struct字段个数: ", num)
	for i := 0; i < num; i++ {
		fmt.Printf("字段 %d,字段值 = %v\n", i, rValue.Field(i))
	}

	methodNumber := r.NumMethod()
	fmt.Println("方法个数：", methodNumber)
	var params []reflect.Value
	params = append(params, reflect.ValueOf("Tom"))
	r.Method(0).Call(params)

}

func main() {
	/*
		编写一个结构体Cal,有两个字段Num1, Num2
		方法GetSub(name string)
		使用反射遍历Cal结构体所有的字段信息
		使用反射机制完成对GetSub的调用，输出形式为
		"Tom 完成了减法运行，8 - 3 = 5"
	*/
	var c Cal = Cal{
		Num1: 33.7,
		Num2: 10,
	}
	TestReflect(&c)

}
