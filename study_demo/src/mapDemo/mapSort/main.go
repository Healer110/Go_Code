package main

import (
	"fmt"
	"sort"
)

func main() {
	// map中的排序演示
	map1 := make(map[int]int)
	map1[10] = 30
	map1[1] = 22
	map1[2] = 36
	map1[3] = 73
	map1[4] = 99
	fmt.Println(map1)

	// 排序步骤
	// 先将map的key放入到切片中
	// 对切片进行排序
	// 遍历切片，然后按照key来输出map的值
	var keySlice []int = make([]int, 0, len(map1))
	for k, _ := range map1 {
		keySlice = append(keySlice, k)
	}

	fmt.Println(keySlice)
	// keySlice[0] = 22

	sort.Ints(keySlice)
	fmt.Println(keySlice)

	for _, v := range keySlice {
		fmt.Println(map1[v])
	}

}
