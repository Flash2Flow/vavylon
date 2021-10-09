package main

import (
	"fmt"
	"log"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
	err_url := [...]string{
		"Error 101:\nBad title ( Empty )",
		"Error 102:\nBad token",
	}

	log.Print("new request")
	query := r.URL.Query()

	title_byte := query.Get("title")
	token_byte := query.Get("token")

	token := string(token_byte[:])
	title := string(title_byte[:])

	if title == "" {
		log.Println(err_url[0])
		fmt.Fprintf(w, err_url[0])
	}

	if token == "" {
		log.Println(err_url[1])
		fmt.Fprintf(w, err_url[1])
	}

	if title == "drop_pass" {
		if token == "cardinal" {

		}

		if token == "quinella" {

		}
	}

	if title == "change_pers" {
		if token == "" {

		}

		if token == "" {

		}
	}

	if title == "" {
		if token == "" {

		}

		if token == "" {

		}
	}
}
