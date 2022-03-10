package service

import (
	"encoding/json"
	"go_programming/chatRoom/common/message"
	"go_programming/chatRoom/common/model"
	"go_programming/chatRoom/common/utils"
	"go_programming/chatRoom/server/dao"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	UserId int
}

func (u *UserProcess) ProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	err = dao.MyUserDao.Login(&model.User{UserId: loginMes.UserId, UserPwd: loginMes.UserPwd})
	if err == nil {
		loginResMes.Code = 200
		u.UserId = loginMes.UserId
		userMgr.NotifyOtherUserOnlineUser(loginMes.UserId)
		userMgr.addOnlineUser(u)
		for id, _ := range userMgr.getOnlineUsers() {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
	} else {
		loginResMes.Code = 403
		loginResMes.Error = err.Error()
	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		return err
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		return err
	}
	tf := utils.Transfer{Conn: u.Conn}
	err = tf.WritePkg(data)
	return
}

func (u *UserProcess) ProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		return err
	}
	var registerResMes message.RegisterResMes
	err = dao.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 400
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 500
			registerResMes.Error = "unknown error"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		return err
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		return err
	}
	tf := utils.Transfer{Conn: u.Conn}
	err = tf.WritePkg(data)
	return

}

func (u *UserProcess) NotifyMeOnlineUser(loginUserId int) (err error) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyMes message.NotifyUserStatusMes
	notifyMes.UserId = loginUserId
	notifyMes.UserStatus = message.UserOnline
	data, err := json.Marshal(notifyMes)
	if err != nil {
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}
	tf := utils.Transfer{Conn: u.Conn}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}
	return
}

func (u *UserProcess) SendMes(mes *message.Message) (err error) {
	data, err := json.Marshal(mes)
	if err!=nil {
		return err
	}
	tf := utils.Transfer{Conn: u.Conn}
	err = tf.WritePkg(data)
	return err
}