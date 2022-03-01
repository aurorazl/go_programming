package model

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
