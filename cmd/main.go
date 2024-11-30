package main

import (
	db "github.com/ChernakovEgor/volley_stats/internal/db"
	server "github.com/ChernakovEgor/volley_stats/internal/server"
	"log"
	"net/http"
)

func main() {
	db := db.CreateDBConnection("dbname= volley_stats_test user=egor")
	db.PingDB()
}

func startServer() {
	mux := http.NewServeMux()
	mux.Handle("/move", http.HandlerFunc(server.RegisterMove))
	log.Fatalln(http.ListenAndServe(":2345", mux))
}
