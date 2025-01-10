package main

import (
	"fmt"
	"reflect"
)

func reflectModify(b interface{}) {
	rValue := reflect.ValueOf(b)
	rValue.Elem().SetInt(202)

}

func main() {
	num := 101
	fmt.Println(num)
	reflectModify(&num)
	fmt.Println(num)
}
