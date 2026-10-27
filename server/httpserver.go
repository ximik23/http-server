package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunServer() {
	for _, link := range api {
		pattern := fmt.Sprintf("%s %s", link.method, link.url)
		http.HandleFunc(pattern, link.handleFunc)
	}

	server := &http.Server{
		Addr:           ":6677",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("server is starting at localhost:6677")

	log.Fatal(server.ListenAndServe())
}
