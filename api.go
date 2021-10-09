package main

import (
	"log"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request)  {
	log.Print("new request")
	//querry := r.URL.Query()
	r.Header.Set("User-Agent", "11")
}
