package process

import "fmt"

var (
	userMgr *UserMsg
)

type UserMsg struct {
	onlineUsers map[int]*UserProcess
}

// 完成对UserMsg初始化工作
func init() {
	userMgr = &UserMsg{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers的添加
func (um *UserMsg) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

// 用户离线，删除
func (um *UserMsg) DeleteOnlineUser(userId int) {
	delete(um.onlineUsers, userId)
}

// 查询,返回当前在线的所有用户
func (um *UserMsg) GetAllOnlineUsers() map[int]*UserProcess {
	return um.onlineUsers
}

// 根据ID，返回一个UserProcess
func (um *UserMsg) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("[%v] 用户没在线", userId)
		return
	}
	return
}
