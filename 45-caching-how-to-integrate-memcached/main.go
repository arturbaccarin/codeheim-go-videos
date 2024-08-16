package main

import (
	"go-memcached/controllers"
	"go-memcached/models"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"

	models.ConnectDatabase()
	models.DBMigrate()

	mux := http.NewServeMux()
	mux.HandleFunc("/blogs/", controllers.BlogShow)

	log.Print("server is listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
