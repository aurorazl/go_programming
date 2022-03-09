package message

import "go_programming/chatRoom/common/model"

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
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
	Code  int
	Error string
}

type RegisterMes struct {
	User model.User `json:"user"`
}

type RegisterResMes struct {
	Code  int
	Error string
}
