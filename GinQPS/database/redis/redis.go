package redis

import (
	"GinQPS/config"
	"GinQPS/model"
	"context"
	"fmt"
	"github.com/fatih/structs"
	"github.com/go-redis/redis/v8"
	"log"
)

var redisConn *redis.Client

var ctx = context.Background()

func init() {
	redisConf := config.Config.DBConf.RedisConf
	// 这个库没有pool的技术（经过测试最大80，不是redis server的上限info clients查看），可以考虑其他库
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

func DeleteKey(key string) {
	_, err := redisConn.Del(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
}

func HMSet(key string, data interface{}) {
	redisConn.HMSet(ctx, key, structs.Map(data))
}

func HSet(key, field, value string) {
	redisConn.HMSet(ctx, key, field, value)
}

func PipelineHSet(models []model.RedisModel) {
	pipeline := redisConn.Pipeline()
	for _, redisModel := range models {
		if redisModel != nil {
			pipeline.HMSet(ctx, redisModel.GetKey(), structs.Map(redisModel))
		}
	}
	_, err := pipeline.Exec(ctx)
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
	val, err := redisConn.HGetAll(ctx, key).Result()
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
