package main

import (
	"log"
	"net/http"

	//"github.com/gorilla/handlers"
	"github.com/jacky-htg/api-news/config"
	"github.com/jacky-htg/api-news/routing"
)

func main() {
	router := routing.NewRouter()
	//allowedHeaders := handlers.AllowedHeaders([]string{"X-Api-Key"})
	//allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	//allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	//err := http.ListenAndServe(config.GetString("server.address"), handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
	err := http.ListenAndServe(config.GetString("server.address"), router)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
