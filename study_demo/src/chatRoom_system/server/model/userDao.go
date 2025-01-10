package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 定义一个结构体，
type UserDao struct {
	pool *redis.Pool
}

// 服务器启动后，就初始化一个UserDao实例，定义为全局变量
var (
	MyUserDao *UserDao
)

// 使用工厂模式，创建一个UserDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 判断用户是否存在，并根据结果返回user 实例或者error
func (dao *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 去redis数据库查询用户信息
	res, err := redis.String(conn.Do("Hget", "users", id))
	if err != nil {
		if err == redis.ErrNil { // 表示没有找到对应的ID
			err = ERROR_USER_NOEXISTS
		}
		return
	}

	user = &User{}
	// 用户存在，check密码，将res 反序列化成user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("[userDao.go/getUserById] 反序列化异常..", err)
		return
	}
	return
}

// 完成登录的校验
func (dao *UserDao) LoginCheck(userId int, userPwd string) (user *User, err error) {
	// 先从连接池中取出连接
	conn := dao.pool.Get()
	defer conn.Close()
	user, err = dao.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 校验密码是否正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 注册功能，完成对redis的新用户添加
func (dao *UserDao) Register(user *User) (err error) {
	conn := dao.pool.Get()
	defer conn.Close()
	_, err = dao.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	// 用户不存在，这里可以进行注册
	data, err := json.Marshal(*user)
	if err != nil {
		return
	}

	_, err = conn.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("[server/userDao.go/Register] error: ", err)
		return
	}

	return
}
