package database

import (
	"context"
	"fmt"
	"github.com/fatih/structs"
	"github.com/go-redis/redis/v8"
	"go_programming/chatRoom/common/config"
	"go_programming/chatRoom/common/model"
	"log"
)

var RedisConn *redis.Client

var Ctx = context.Background()

func init() {
	redisConf := config.Config.DBConf.RedisConf
	fmt.Println(redisConf)
	// 这个库底层有connPool技术（经过测试最大80，不是redis server的上限info clients查看），可以考虑其他库
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", redisConf.IP, redisConf.Port),
		Password:     redisConf.Password, // no password set
		DB:           redisConf.Database, // use default DB
		PoolSize:     20,                 //Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
		MinIdleConns: 20,
	})
	RedisConn = rdb
}

func Scan(pattern string) []string {
	key_list := []string{}
	var cursor uint64
	var keys []string
	var err error
	for {
		keys, cursor, err = RedisConn.Scan(Ctx, cursor, pattern, 1000).Result()
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
		keys, cursor, err = RedisConn.SScan(Ctx, key, cursor, "*", 1000).Result()
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
	result, err := RedisConn.Exists(Ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	if result == 1 {
		return true
	}
	return false
}

func DeleteKey(key string) {
	_, err := RedisConn.Del(Ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
}

func HMSet(key string, data interface{}) {
	RedisConn.HMSet(Ctx, key, structs.Map(data))
}

func HSet(key, field, value string) {
	RedisConn.HMSet(Ctx, key, field, value)
}

func PipelineHSet(models []model.RedisModel) {
	pipeline := RedisConn.Pipeline()
	for _, redisModel := range models {
		if redisModel != nil {
			pipeline.HMSet(Ctx, redisModel.GetKey(), structs.Map(redisModel))
		}
	}
	_, err := pipeline.Exec(Ctx)
	if err != nil {
		panic(err)
	}
	//for _, cmd := range cmds {
	//	fmt.Println(cmd.(*redis.StringCmd).Val())
	//}
}

func PipelineHSetInBatch(models []model.RedisModel, batchSize int) {
	length := len(models)
	var i int
	for ; i <= length/batchSize; i++ {
		start := i * batchSize
		if start >= length {
			return
		}
		end := (i + 1) * batchSize
		if end >= length {
			end = length - 1
		}
		PipelineHSet(models[i*batchSize : (i+1)*batchSize])
	}
}

func HGetAll(key string) map[string]string {
	val, err := RedisConn.HGetAll(Ctx, key).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
		return nil
	case err != nil:
		fmt.Println("Get failed", err)
		return nil
	default:
		return val
	}
}
