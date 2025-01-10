package main

import "fmt"

func InsertSort(arr *[7]int) {
	// 完成第一次，给第二个元素找到合适的位置并插入
	// insertVal := arr[1]
	// insertIndex := 0

	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		// 上面的循环中，最终insertIndex做了insertIndex--, 所以这里要+1,如果最终的位置没有变化，就不做处理了
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次的排序结果：%v\n", i+1, *arr)
	}
}

func main() {
	arr := [7]int{23, 0, 12, 56, 34, -1, 55}
	fmt.Println("排序前：", arr)
	InsertSort(&arr)
	fmt.Println("排序后：", arr)
}
