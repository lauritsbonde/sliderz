package main

import (
	"fmt"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Welcome to the home page!")
}

func hello(w http.ResponseWriter, req *http.Request){
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(1 * time.Second):
		fmt.Fprintf(w, "Hello, World!")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func headers(w http.ResponseWriter, req *http.Request){
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}