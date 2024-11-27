package main

import (
	// "bufio"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/move", http.HandlerFunc(registerMove))
	log.Fatalln(http.ListenAndServe(":2345", mux))
}

func registerMove(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		body := make([]byte, 1000)
		r.Body.Read(body)
		log.Println("OPTIONS:", string(body))

		w.Header().Add("Access-Control-Allow-Origin:", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With,content-type")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(200)
	}
	if r.Method == "POST" {
		log.Println("POST")
	}
	// w.WriteHeader(201)
}
