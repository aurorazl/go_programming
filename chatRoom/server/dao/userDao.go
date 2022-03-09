package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go_programming/chatRoom/common/model"
	"strconv"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Client
	ctx  context.Context
}

func NewUserDao(pool *redis.Client, ctx context.Context) *UserDao {
	return &UserDao{pool, ctx}
}

func (u *UserDao) GetUserById(id int) (user model.User, err error) {
	res, err := u.pool.Get(u.ctx, strconv.Itoa(id)).Result()
	if err != nil {
		if err == redis.Nil {
			err = model.ERROR_USER_NOTEXISTS
		}
		fmt.Println("redis get error", err)
		return
	}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		return
	}
	return
}

func (u *UserDao) Login(user *model.User) (err error) {
	userInRedis, err := u.GetUserById(user.UserId)
	if err != nil {
		return
	}
	if userInRedis.UserPwd != user.UserPwd {
		return model.ERROR_PWD
	}
	return
}

func (u *UserDao) Register(user *model.User) (err error) {
	_, err = u.GetUserById(user.UserId)
	if err == nil {
		err = model.ERROR_USER_TEXISTS
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = u.pool.Set(u.ctx, strconv.Itoa(user.UserId), string(data), redis.KeepTTL).Result()
	if err != nil {
		return err
	}
	return
}
