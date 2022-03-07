package main

import (
	"GinQPS/model"
	"GinQPS/service"
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
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
			Favorite: favorites[rand.Intn(4)] + uuid.New().String()[:5],
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
			generateUserInBatch(500000, i)
			<-queue
			w.Done()
		}(i)
	}
	time.Sleep(10000)
	w.Wait()
	fmt.Println("insert done")
}

func contextDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(ctx context.Context, num int) {
			select {
			case <-ctx.Done():
				break
			default:
				fmt.Println("goroutine ", num, " waiting for signal")
				time.Sleep(time.Second)
			}
		}(ctx, i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("begin cancel")
	cancel()
}

func contextTimeoutDemo() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	respChan := make(chan int)
	for i := 0; i < 5; i++ {
		go func(ctx context.Context, num int) {
			time.Sleep(time.Second * 1)
			fmt.Println("task finish!")
			respChan <- 1
		}(ctx, i)
	}
	select {
	case <-ctx.Done():
		fmt.Println("timeout") // no active cancel
	case <-respChan:
		fmt.Println("success")
	}
}

func selectDemo() {
	a := make(chan int)
	close(a)
	var i int
	select {
	case i = <-a:
		fmt.Println(i)
	default:
		fmt.Println("default")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	//batchGenerateUser()
}
