package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println("hello world")
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
