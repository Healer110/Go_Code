package main

import "fmt"

/*
	快速排序：先找一个位置的数当做中间值，任意位置都可以
	将数组的中其他值跟中间值进行比较，比中间值大的值放到右边
	比中间值小的值放到坐标，递归后，就可以全部排出来
*/
func QuickSort(left int, right int, arr *[6]int) {
	temp := arr[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && arr[j] >= temp {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}

		for i <= p && arr[i] <= temp {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = temp
	if p-left > 1 {
		QuickSort(left, p-1, arr)
	}
	if right-p > 1 {
		QuickSort(p+1, right, arr)
	}

}

func main() {
	arr := [6]int{-9, 78, 0, 23, -567, -70}
	fmt.Println("排序前：", arr)
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("排序后：", arr)
}
