package service

import (
	"GinQPS/database/mysql"
	"GinQPS/database/redis"
	"GinQPS/model"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func GetUserById(id int64) model.User {
	var user model.User
	mysql.Db.First(&user, id)
	return user
}

func GetUserByName(name string) model.User {
	var user model.User
	mysql.Db.Where("name=?", name).Find(&user)
	return user
}

func GetUserByIdOnRedis(id int64) model.User {
	var user model.User
	result := redis.HGetAll(fmt.Sprintf("1_%d", id))
	marshal, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(marshal, &user)
	if err != nil {
		log.Println(err)
	}
	user.Id, err = strconv.ParseInt(result["Id"], 10, 64)
	if err != nil {
		log.Println(err)
	}
	age, err := strconv.ParseInt(result["Age"], 10, 8)
	user.Age = int8(age)
	if err != nil {
		log.Println(err)
	}
	salary, err := strconv.ParseInt(result["Salary"], 10, 32)
	user.Salary = int32(salary)
	if err != nil {
		log.Println(err)
	}
	return user
}

func InsertUser(user model.User) {
	mysql.Db.Create(&user)
}

func BatchInsertUser(users []model.User, batchSize int) {
	mysql.Db.CreateInBatches(&users, batchSize)
}
