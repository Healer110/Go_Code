package model

type User struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"username"`
	UserStatus int    `json:"userStatus"`
}
