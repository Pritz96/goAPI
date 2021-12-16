package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("API has started on port 8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", mux)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		tm := time.Now().Format(time.RFC1123)
		w.Write([]byte("The time is " + tm))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("HTTP request method not allowed"))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("HTTP request method not allowed"))
	}
}
