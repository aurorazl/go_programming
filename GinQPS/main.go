package main

import (
	"GinQPS/config"
	"GinQPS/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"log"
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

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
