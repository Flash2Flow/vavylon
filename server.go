package main

import (
	"log"
	"net/http"
	//"os"
)

func server() {
	//port := os.Getenv("PORT")
	log.Println("server start with port: "/*+port*/)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080"/*+port*/, nil)
}
