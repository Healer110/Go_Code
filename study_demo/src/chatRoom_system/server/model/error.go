package model

import "errors"

// 自定义错误消息
var (
	ERROR_USER_NOEXISTS = errors.New("用户不存在...请先注册...")
	ERROR_USER_EXISTS   = errors.New("用户已存在...")
	ERROR_USER_PWD      = errors.New("密码不正确...")
)
