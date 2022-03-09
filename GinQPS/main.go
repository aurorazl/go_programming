package main

import (
	"GinQPS/config"
	"GinQPS/controller"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := config.Config.ServerConfig.Port

	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	controller.AddGroupUser(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// hystrix config
	hystrix.ConfigureCommand("default_command", hystrix.CommandConfig{
		Timeout:               3000,
		MaxConcurrentRequests: 1000000,
		ErrorPercentThreshold: 25,
		RequestVolumeThreshold: 10, // 统计窗口10s内的请求数量，达到这个请求数量后才去判断是否要开启熔断
		SleepWindow:           2000,
	})

	// hystrix metric upload
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "8081"), hystrixStreamHandler)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
