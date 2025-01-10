package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("====== start ======")
	fmt.Println(s)
	fmt.Println("====== end ======")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	rType := reflect.TypeOf(a)
	rValue := reflect.ValueOf(a)
	rKind := rValue.Kind()
	if rKind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := rValue.NumField()
	fmt.Printf("struct has %v field \n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, rValue.Field(i))
		// 获取struct的标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := rType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	numOfMethod := rValue.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	rValue.Method(1).Call(nil)
	fmt.Println("method by name...")
	rValue.MethodByName("Print").Call(nil)

	// 调用结构体的第1个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rValue.Method(0).Call(params)
	fmt.Println("res =", res[0].Int())

}

func main() {
	/*
		使用反射遍历结构体字段，调用结构体的方法，并获取结构体标签的值
		Method(), Call()
	*/
	var a Monster = Monster{
		Name:  "乌龟",
		Age:   110,
		Score: 88,
		Sex:   "Male",
	}
	TestStruct(a)
}
