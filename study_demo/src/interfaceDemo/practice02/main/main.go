package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 使用接口，实现对hero切片的排序
// type MySort interface {
// 	MSort()
// }

// type Slice []int

// func (slice *Slice) MSort() {
// 	for i := 0; i < len(*slice)-1; i++ {
// 		for j := 0; j < len(*slice)-1-i; j++ {
// 			if (*slice)[j] > (*slice)[j+1] {
// 				tmp := 0
// 				tmp = (*slice)[j]
// 				(*slice)[j] = (*slice)[j+1]
// 				(*slice)[j+1] = tmp
// 			}
// 		}
// 	}
// }

// func main() {
// 	var ms Slice = []int{1, 5, 3, 99, 66}
// 	fmt.Println("Before Sort:", ms)
// 	var s MySort = &ms
// 	s.MSort()
// 	fmt.Println("After Sort:", ms)
// }

type Hero struct {
	Name string
	Age  int
}

// 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// 决定按照什么标准进行排序：按照年龄进行排序
// > : 逆排序； < 正排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	// var tmp = hs[i]
	// hs[i] = hs[j]
	// hs[j] = tmp
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var hs HeroSlice
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero%d", i),
			Age:  rand.Intn(100) + 1,
		}
		hs = append(hs, hero)
	}
	fmt.Println("排序前：")
	for _, v := range hs {
		fmt.Printf("name: %v, age: %v \n", v.Name, v.Age)
	}
	sort.Sort(hs)
	fmt.Println("排序后：")
	for _, v := range hs {
		fmt.Printf("name: %v, age: %v \n", v.Name, v.Age)
	}
}
