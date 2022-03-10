package service

import (
	"errors"
	"fmt"
)

var (
	userMgr UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = UserMgr{make(map[int]*UserProcess, 1024)}
}

func (u *UserMgr)addOnlineUser(up *UserProcess) {
	u.onlineUsers[up.UserId] = up
}

func (u *UserMgr)deleteOnlineUser(up *UserProcess) {
	delete(u.onlineUsers,up.UserId)
}

func (u *UserMgr) getOnlineUsers() map[int]*UserProcess {
	return u.onlineUsers
}

func (u *UserMgr) getOnlineUserById(userId int)  (userProcess *UserProcess, err error) {
	userProcess,ok := u.onlineUsers[userId]
	if ! ok {
		err = errors.New(fmt.Sprintf("user %d not exists", userId))
		return
	}
	return
}

func (u *UserMgr) NotifyOtherUserOnlineUser(loginUserId int) {
	for id,userProcess := range u.onlineUsers {
		if id == loginUserId {
			continue
		}
		userProcess.NotifyMeOnlineUser(loginUserId)
	}
}

