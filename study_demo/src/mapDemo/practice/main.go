package main

import (
	"fmt"
	"strconv"
)

func main() {
	userMap := make(map[string]map[string]string)
	for i := 0; i < 3; i++ {
		var tmp string = strconv.Itoa(i)
		user := make(map[string]string, 2)
		user["nickname"] = "May" + tmp
		user["pwd"] = "666666"
		userMap["user"+tmp] = user
	}

	fmt.Printf("before run: %v \n", userMap)
	modifyUser(userMap, "user0")
	fmt.Printf("after run: %v \n", userMap)

}

/*
使用map[string]map[string]string的map类型
key: 标识用户名，是唯一的，不可以重复
如果某个用户名存在，就将其密码修改为“888888”，如果不存在就增加这个用户信息（包括昵称nickname和密码pwd）
编写一个函数modifyUser(users map[string]map[string]string, name string)完成上述功能
*/
func modifyUser(users map[string]map[string]string, name string) {
	// for userId, v := range users {
	// 	if name == userId {
	// 		v["pwd"] = "888888"
	// 		return
	// 	}
	// }
	// // 用户名不存在，新增用户名
	// var user map[string]string = make(map[string]string)
	// user["nickname"] = "May"
	// user["pwd"] = "000000"
	// users[name] = user

	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		var user map[string]string = make(map[string]string)
		user["nickname"] = "May"
		user["pwd"] = "000000"
		users[name] = user
	}
}
