package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	demo01()
	demo02()
	demo03()
}

// 序列化时，struct里字段首字母必须大写
type stu struct {
	Name   string
	Age    int
	Gender string
}

// 结构体序列化和反序列化
func demo01() {
	stu1 := stu{
		Name:   "Tom",
		Age:    28,
		Gender: "男",
	}
	// stu2 := stu{"Jack", 33, "男"}
	// stu3 := stu{"Mary", 23, "女"}
	data, err := json.Marshal(&stu1)
	if err != nil {
		fmt.Println("序列化失败: ", err)
	} else {
		fmt.Println(string(data))
	}

	// 反序列化
	var respStu stu
	res := string(data)
	json.Unmarshal([]byte(res), &respStu)
	fmt.Println(respStu)

}

// 将map序列化
func demo02() {
	// 定义一个map
	var a map[string]interface{} = make(map[string]interface{})
	a["name"] = "zzz"
	a["age"] = 30
	a["address"] = "山~~"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化异常...")
	} else {
		fmt.Println(string(data))
	}
}

// 将切片序列化
func demo03() {
	var slice []map[string]interface{}

	var m1 map[string]interface{} = make(map[string]interface{})
	m1["name"] = "Jack"
	m1["age"] = 47
	m1["address"] = "北京"

	slice = append(slice, m1)

	var m2 map[string]interface{} = make(map[string]interface{})
	m2["name"] = "Tom"
	m2["age"] = 28
	m2["address"] = [2]string{"北京", "美国"}

	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化异常...")
	} else {
		fmt.Println(string(data))
	}
}
