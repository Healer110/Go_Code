package main

import (
	"fmt"
	"reflect"
)

func reflectConvert(b interface{}) {
	rValue := reflect.ValueOf(b)

	rType := rValue.Type()
	rKind := rValue.Kind()
	fmt.Println(rType, rKind)

	iValue := rValue.Interface().(float64) + 3
	fmt.Printf("iValue = %v, iValue type = %T \n", iValue, iValue)

}

func reflectModifyStr(s interface{}) {
	fs := reflect.ValueOf(s)
	fs.Elem().SetString("Jack")

}

func main() {
	/*
		给你一个变量 var v flot64 = 1.2 请使用反射来得到它的reflect.Value, 然后获取对
		赢得Type, Kind和值，并将reflect.Value转换成空接口，再将空接口转换成float64
	*/

	var v float64 = 1.2
	reflectConvert(v)

	// 修改string的值
	var str string = "Tom"
	reflectModifyStr(&str)
	fmt.Println(str)

}
