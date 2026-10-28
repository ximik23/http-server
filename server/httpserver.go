package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunServer() {
	mux := http.NewServeMux()
	for _, link := range api {
		pattern := fmt.Sprintf("%s /api/v1%s", link.method, link.url)
		fmt.Println(pattern)
		mux.HandleFunc(pattern, link.handleFunc)
	}

	server := &http.Server{
		Addr:           ":6677",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("server is running at localhost:6677")

	log.Fatal(server.ListenAndServe())
}
