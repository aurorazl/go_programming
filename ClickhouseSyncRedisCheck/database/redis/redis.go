package redis

import (
	"ClickhouseSyncRedisCheck/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var redisConn *redis.Client

var ctx = context.Background()

func init() {
	redisConf := config.Config.DBConf.RedisConf
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConf.IP, redisConf.Port),
		Password: redisConf.Password, // no password set
		DB:       redisConf.Database, // use default DB
	})
	redisConn = rdb
}

func Scan(pattern string) []string {
	key_list := []string{}
	var cursor uint64
	var keys []string
	var err error
	for {
		keys, cursor, err = redisConn.Scan(ctx, cursor, pattern, 1000).Result()
		if err != nil {
			log.Fatal(err)
		}
		key_list = append(key_list, keys...)
		if cursor == 0 {
			break
		}
	}
	return key_list
}

func SScan(key string) []string {
	key_list := []string{}
	var cursor uint64
	var keys []string
	var err error
	for {
		keys, cursor, err = redisConn.SScan(ctx, key, cursor, "*", 1000).Result()
		if err != nil {
			log.Fatal(err)
		}
		key_list = append(key_list, keys...)
		if cursor == 0 {
			break
		}
	}
	return key_list
}

func KeyExists(key string) bool {
	result, err := redisConn.Exists(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	if result == 1 {
		return true
	}
	return false
}
