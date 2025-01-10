package main

import (
	"testing"
)

// 编写测试用例, 使用testing测试框架
func TestAddUpper(t *testing.T) {
	// 调用
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("异常...\n")
	}
	// fmt.Println("TestEncrypt = ", res)

	t.Logf("执行正确\n")
}
