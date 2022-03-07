package main

import (
	"GinQPS/database/mysql"
	"GinQPS/database/redis"
	"GinQPS/model"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func scanMysqlToRedis() {
	var users []model.User
	mysql.Db.Find(&users)
	// limit offset

	var w sync.WaitGroup
	n := len(users)
	fmt.Println("scan done, length is ", n)
	num := n / 100000
	for i := 0; i <= num; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			for index := i * 100000; index < (i+1)*100000 && index < n; index++ {
				if index%1000 == 0 {
					fmt.Println("goroutine:"+strconv.Itoa(i)+" current sync index is ", index)
				}
				redis.HMSet(fmt.Sprintf("1_%d", users[index].Id), users[index])
			}
		}(i)
	}
	time.Sleep(2000)
	w.Wait()
	fmt.Println("sync done")
}

func batchScanMysqlToRedis() {
	var users []model.User
	mysql.Db.Find(&users)
	// limit offset
	var models []model.RedisModel
	// 通过append转化struct为接口
	for _, user := range users {
		models = append(models, user)
	}
	var w sync.WaitGroup
	totalLength := len(users)
	fmt.Println("scan done, length is ", totalLength)
	num := totalLength / 100000
	for i := 0; i <= num; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			start := i * 100000
			if start >= totalLength {
				return
			}
			end := (i + 1) * 100000
			if end >= totalLength {
				end = totalLength - 1
			}
			redis.PipelineHSetInBatch(models[start:end], 3000)
		}(i)
	}
	time.Sleep(2000)
	w.Wait()
	fmt.Println("sync done")
}

//func main() {
//	batchScanMysqlToRedis()
//}
