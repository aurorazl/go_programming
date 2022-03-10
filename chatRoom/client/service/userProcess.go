package service

import (
	"encoding/json"
	"fmt"
	"go_programming/chatRoom/common/message"
	"go_programming/chatRoom/common/model"
	"go_programming/chatRoom/common/utils"
	"net"
)

type UserProcess struct {
}

func (u *UserProcess) Login(userId int, userPwd string) (err error) {
	fmt.Printf("user is %d user password %s\n", userId, userPwd)
	dial, err := net.Dial("tcp", "0.0.0.0:8899")
	tf := utils.Transfer{Conn: dial}
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dial.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = tf.WritePkg(data)
	if err != nil {
		return err
	}
	mes, err = tf.ReadPkg()
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("login success")
		for _, onlineUserId := range loginResMes.UserIds {
			userMgr.onlineUsers[onlineUserId] = &model.User{UserId: onlineUserId, UserStatus: message.UserOnline}
		}
		userMgr.smsProcess = &SmsProcess{userId, &tf}
		go serverProcessMes(dial)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println("login error", loginResMes.Error)
	}
	return nil
}

func (u *UserProcess) Register(userId int, userPwd, userName string) (err error) {
	dial, err := net.Dial("tcp", "0.0.0.0:8899")
	tf := utils.Transfer{Conn: dial}
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dial.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = tf.WritePkg(data)
	if err != nil {
		return err
	}
	mes, err = tf.ReadPkg()
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("register success")
	} else {
		fmt.Println("register error", registerResMes.Error)
	}
	return nil
}
