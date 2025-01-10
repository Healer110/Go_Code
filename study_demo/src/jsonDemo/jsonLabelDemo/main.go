package main

import (
	"encoding/json"
	"fmt"
)

// 序列化时，struct里字段首字母必须大写,
// 使用alias时冒号后不能有空格
type stu struct {
	Name   string `json:"stu_name"`
	Age    int    `json:"stu_age"`
	Gender string `json:"stu_gender`
}

func main() {
	stu1 := stu{
		Name:   "Tom",
		Age:    28,
		Gender: "男",
	}

	data, err := json.Marshal(&stu1)
	if err != nil {
		fmt.Println("序列化失败: ", err)
	} else {
		fmt.Println(string(data))
	}

}
