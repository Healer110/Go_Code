package main

import (
	"fmt"
	"redisDemo/utils"

	"github.com/garyburd/redigo/redis"
)

// 操作Redis的示例
func operateRedisDemo() {
	// 通过Go向Redis写入数据，读取数据
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("Connect redis error...", err)
		return
	}
	fmt.Println("connect successfully...", conn)

	// 写入数据
	_, err = conn.Do("set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set error...", err)
		return
	}
	fmt.Println("写入成功...")

	// 读取数据
	// res, err := conn.Do("get", "name")
	res, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get error...", err)
		return
	}
	// 返回的res是空接口类型，这里需要类型断言
	// fmt.Printf("读取成功...name = %v \n", string(res.([]uint8)))
	// 使用Redis的方法进行转换
	fmt.Printf("读取成功...name = %v \n", res)

}

// 操作hash
func operateHash() {
	// 通过Go向Redis写入数据，读取数据
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("Connect redis error...", err)
		return
	}
	fmt.Println("connect successfully...", conn)

	// 写入数据
	_, err = conn.Do("Hset", "user01", "name", "john")
	if err != nil {
		fmt.Println("hset error...", err)
		return
	}

	_, err = conn.Do("Hset", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset error...", err)
		return
	}
	fmt.Println("写入成功...")

	// 读取数据
	// res, err := conn.Do("get", "name")
	res, err := redis.String(conn.Do("Hget", "user01", "name"))
	if err != nil {
		fmt.Println("hget error...", err)
		return
	}

	age, err := redis.Int(conn.Do("Hget", "user01", "age"))
	if err != nil {
		fmt.Println("hget error...", err)
		return
	}

	// 返回的res是空接口类型，这里需要类型断言
	// fmt.Printf("读取成功...name = %v \n", string(res.([]uint8)))
	// 使用Redis的方法进行转换
	fmt.Printf("读取成功...name = %v, age = %v \n", res, age)
}

// practice
// Monster信息（那么，age，skill)
// 通过终端输入三个Monster的信息，使用Golang操作redis，存放到redis中（比如使用hash数据类型）
// 编程，遍历出所有的Monster信息，并显示在终端
// 可以先比遍历keys monster*，然后遍历Monster中的字段
func rangeMonsterDemo() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("connect error: ", err)
		return
	}

	monsters, err := conn.Do("keys", "monster*")
	if err != nil {
		fmt.Println("get monster keys error: ", err)
		return
	}
	fmt.Println(monsters)
	fmt.Printf("%T \n", monsters)
	for _, v := range monsters.([]interface{}) {
		m, _ := redis.String(v, nil)
		fmt.Println(m)

		// 遍历各个Monster
		info, err := redis.Strings(conn.Do("Hmget", m, "name", "age", "skill"))
		if err != nil {
			fmt.Println("获取失败：", err)
		}
		// fmt.Println(info)
		for _, v := range info {
			fmt.Println(v)
		}

	}

}

func main() {
	// Go mod机制导入自定义包
	utils.Hello()
	fmt.Println("main func....")

	// operateRedisDemo()
	// operateHash()
	rangeMonsterDemo()

}
