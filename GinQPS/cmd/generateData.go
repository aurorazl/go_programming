package main

import (
	"GinQPS/model"
	"GinQPS/service"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"sync"
)

var favorites = []string{"play", "run", "eat", "sleep"}

func generateUser(num int64, goroutineNum int) {
	for i := 0; int64(i) < num; i++ {
		if i%1000 == 0 {
			fmt.Println("goroutine: ", goroutineNum, " current loop is ", i)
		}
		user := model.User{
			Name:     uuid.New().String(),
			Age:      int8(rand.Intn(100)),
			Favorite: favorites[rand.Intn(4)],
			Salary:   rand.Int31n(10000)}
		service.InsertUser(user)
	}
}

//速度提升很大
func generateUserInBatch(num int64, goroutineNum int) {
	var users []model.User
	for i := 0; int64(i) < num; i++ {
		if i%1000 == 0 {
			fmt.Println("goroutine: ", goroutineNum, " current loop is ", i)
		}
		users = append(users, model.User{
			Name:     uuid.New().String(),
			Age:      int8(rand.Intn(100)),
			Favorite: favorites[rand.Intn(4)],
			Salary:   rand.Int31n(10000)})
	}
	service.BatchInsertUser(users, 10000)
}

func batchGenerateUser() {
	var w sync.WaitGroup
	queue := make(chan interface{}, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			queue <- struct{}{}
			w.Add(1)
			generateUserInBatch(100000, i)
			<-queue
			w.Done()
		}(i)
	}
	w.Wait()
	fmt.Println("insert done")
}

func main() {
	batchGenerateUser()
}
