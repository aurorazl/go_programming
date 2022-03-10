package service

import (
	"encoding/json"
	"fmt"
	"go_programming/chatRoom/common/message"
	"go_programming/chatRoom/common/model"
)

var (
	userMgr UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*model.User
	smsProcess  *SmsProcess
}

func init() {
	userMgr = UserMgr{onlineUsers: make(map[int]*model.User, 1024)}
}

func (u *UserMgr) updateOnlineUserStatus(mes *message.Message) (err error) {
	var notifyMes message.NotifyUserStatusMes
	err = json.Unmarshal([]byte(mes.Data), &notifyMes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	user, ok := userMgr.onlineUsers[notifyMes.UserId]
	if !ok {
		user = &model.User{UserId: notifyMes.UserId}

	}
	user.UserStatus = notifyMes.UserStatus
	userMgr.onlineUsers[notifyMes.UserId] = user
	return err
}

func (u *UserMgr) showAllOnlineUsers() {
	for userId, _ := range u.onlineUsers {
		fmt.Println("user id:", userId)
	}
}
