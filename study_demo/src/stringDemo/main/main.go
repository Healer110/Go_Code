package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// for_range_demo()
	// runeDemo()
	numberTransform()
}

func runeDemo() {
	fmt.Println("rune Demo")
	str := "hello 北京，我来了"
	tmp := []rune(str)
	for i := 0; i < len(tmp); i++ {
		fmt.Printf("idx=%d\tvalue=%c\n", i, tmp[i])
	}
}

func for_range_demo() {
	str := "hello 北京，我来了"

	for i := 0; i < len(str); i++ {
		fmt.Printf("idx=%d\tvalue=%c\n", i, str[i])
	}

	fmt.Println("for range Demo")
	for idx, val := range str {
		fmt.Printf("idx=%d\tvalue=%c\n", idx, val)
	}
}

func numberTransform() {
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println("num =", num)
	}

	str := strconv.Itoa(12345)
	fmt.Println("str =", str)

	// 按照指定的标识，拆分字符串
	str = "Hello, beijng, 你好, beijing, 再见, beijing"
	strArr := strings.Split(str, ",")
	fmt.Println("str list:", strArr)
	fmt.Println(strings.Index(str, "H"))
	fmt.Println(strings.LastIndex(str, "b"))
	fmt.Println(strings.ToLower("Go Go go"))
	fmt.Println(strings.ToUpper("Go Go go"))
	fmt.Println(strings.TrimSpace("  Go Go go  "))
	fmt.Println(strings.TrimLeft("  Go Go go aabc", " "))
	fmt.Println(strings.TrimRight("  Go Go go aabc", " abc"))
	fmt.Println(strings.Trim("aabcGo Go goaabc", "abc"))

}
