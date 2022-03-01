package service

import (
	"GinQPS/database/mysql"
	"GinQPS/model"
)

func GetUserById(id int64) model.User {
	var user model.User
	mysql.Db.First(&user, id)
	return user
}

func InsertUser(user model.User) {
	mysql.Db.Create(&user)
}

func BatchInsertUser(users []model.User, batchSize int) {
	mysql.Db.CreateInBatches(&users, batchSize)
}
