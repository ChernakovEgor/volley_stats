package server

import (
	"log"
	"net/http"
)

func RegisterMove(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		log.Println("OPTIONS request")
		w.WriteHeader(204)
	}

	if r.Method == "POST" {
		body := make([]byte, 1000)
		r.Body.Read(body)
		log.Println("POST:", string(body))
		w.WriteHeader(201)
	}

	if r.Method == "GET" {
		body := make([]byte, 1000)
		r.Body.Read(body)
		log.Println("GET:", string(body))
		w.WriteHeader(200)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH")
	(*w).Header().Add("Access-Control-Allow-Headers", "X-Requested-With,content-type")
	(*w).Header().Add("Access-Control-Allow-Credentials", "true")
}
