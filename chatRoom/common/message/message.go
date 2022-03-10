package message

import "go_programming/chatRoom/common/model"

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string
	Data string
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code    int
	Error   string
	UserIds []int
}

type RegisterMes struct {
	User model.User `json:"user"`
}

type RegisterResMes struct {
	Code  int
	Error string
}

type NotifyUserStatusMes struct {
	UserId     int `json:"userId"`
	UserStatus int `json:"userStatus"`
}

type SmsMes struct {
	Content string `json:"content"`
	// 发送人
	model.User
}
