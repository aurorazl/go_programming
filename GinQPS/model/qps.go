package model

import "fmt"

type User struct {
	Id       int64
	Name     string
	Age      int8
	Favorite string
	Salary   int32
}

func (User) TableName() string {
	return "qps.user"
}

func (u User) GetKey() string {
	return fmt.Sprintf("1_%d", u.Id)
}
