package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func httpServer() {
	fmt.Printf("listen on 8081")
	s := grpc.NewServer()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		s.ServeHTTP(writer, request)
	})

	httpServer := &http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	err := httpServer.ListenAndServeTLS("./identity/server.crt", "./identity/server_no_password.key")
	if err != nil {
		log.Fatal(err)
	}

}
