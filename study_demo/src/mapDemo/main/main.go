package main

import (
	"fmt"
	"strconv"
)

func main() {
	// demo01()
	// demo02()
	// demo03()
	sliceOfMapDemo()
}

func demo01() {
	// 定义一个map
	var a map[string]string
	a = make(map[string]string, 5)

	a["no1"] = "Jone"
	a["no2"] = "vvv"
	a["no1"] = "abc"

	fmt.Println(a)
}

// 定义一个map存放学生信息，每个学生有姓名跟性别以及班级信息
func demo02() {
	var studentInfo = map[string]map[string]string{
		"John":  {"gender": "male", "class": "no1"},
		"Lixy":  {"gender": "female", "class": "no2"},
		"Zotia": {"gender": "male", "class": "no1"},
	}

	fmt.Println(studentInfo)
}

// map的增删改查
func demo03() {
	var mapA = make(map[string]string, 5)
	mapA["zzz"] = "hhh"
	mapA["www"] = "hhh~"
	mapA["yyy"] = "hhh~~~"
	delete(mapA, "zzzz")
	fmt.Println(mapA)

	// map的查找
	val, flag := mapA["zzz1"]
	if flag {
		fmt.Println(val)
	} else {
		fmt.Println(flag)
	}

	// map的遍历
	for k, v := range mapA {
		fmt.Printf("%v = %v\n", k, v)
	}
}

func sliceOfMapDemo() {
	// map切片演示
	// 使用切片记录人员信息，每个人的信息放置在一个map数据结构里，map数据包括ame and age
	var personInfo []map[string]string = make([]map[string]string, 5)

	for i := 0; i < len(personInfo); i++ {
		person := make(map[string]string, 2)
		str := strconv.Itoa(i)
		person["name"+str] = "Jp" + str
		person["age"+str] = "24"
		// 使用下面的方式直接赋值，当数值超出切片定义的容量后，就会报错
		// 想实现切片的动态增加，使用append内置方法新增元素
		personInfo[i] = person
	}

	person := make(map[string]string, 2)
	person["name"] = "Jp"
	person["age"] = "24"
	personInfo = append(personInfo, person)

	for _, v := range personInfo {
		fmt.Println(v)
	}
}
