package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
	 ErrUrl := [...]string{
		"Error 101:\nBad title ( Empty )",
		"Error 102:\nBad token",
	}

	log.Print("new request")
	query := r.URL.Query()

	TitleByte := query.Get("title")
	TokenByte := query.Get("token")
	EmailByte := query.Get("email")
	UserKeyByte := query.Get("key")
	PasswordByte := query.Get("password")
	LoginByte := query.Get("login")


	title := string(TitleByte[:])		//err id 1
	token := string(TokenByte[:])		//err id 2
	email := string(EmailByte[:])		//err id 3
	userkey := string(UserKeyByte[:])	//err id 4
	password := string(PasswordByte[:])	//err id 5
	login := string(LoginByte[:])		//err id 6


	if title == "" {
		log.Println(ErrUrl[0])
		fmt.Fprintf(w, ErrUrl[0])
	}

	if token == "" {
		log.Println(ErrUrl[1])
		fmt.Fprintf(w, ErrUrl[1])
	}

	if title == "drop_pass" {
		if token != "cardinal" {
			OtherErrors := [...]string{
				"Error 202:\rIncorrect Token Cardinal",
			}
			log.Println(OtherErrors[0])
			fmt.Fprintf(w, OtherErrors[0])
		}else{
			if email == "" {
				ErrUrl := [...]string{
					"Error 103:\rIncorrect Email ( Empty )",
				}
				log.Println(ErrUrl[0])
				fmt.Fprintf(w, ErrUrl[0])
			}else{
				if userkey == "" {
					ErrUrl := [...]string{
						"Error 104:\rIncorrect User Key ( Empty )",
					}
					log.Println(ErrUrl[0])
					fmt.Fprintf(w, ErrUrl[0])
				}else{
					//http req cardinal

					resp, err := http.Get("https://vavylon1cardinal.herokuapp.com/api?title=drop_pass&token=cardinal&email=email&userkey=userkey")
					if err != nil {
						print(err)
					}
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						print(err)
					}
					//fmt.Print(string(body))
					var res map[string]interface{}
					json.Unmarshal([]byte(body), &res)
					if res["Status"] == "OK"{
						log.Println("Nice request")
					}

					if res["Status"] == "Err, bad token"{
						OtherErrors := [...]string{
							"Error 202:\rIncorrect Token Cardinal",
						}
						log.Println(OtherErrors[0])
						fmt.Fprintf(w, OtherErrors[0])
					}

					if res["Status"] == "Err, bad user key"{
						OtherErrors := [...]string{
							"Error 204:\rIncorrect User Key",
						}
						log.Println(OtherErrors[0])
						fmt.Fprintf(w, OtherErrors[0])
					}

				}
				}
			}
		}

	if title == "auth" {
		if token != "cardinal" {
			OtherErrors := [...]string{
				"Error 202:\rIncorrect Token Cardinal",
			}
			log.Println(OtherErrors[0])
			fmt.Fprintf(w, OtherErrors[0])
		}else{
			//send req cardinal
			if login == "" {
				ErrUrl := [...]string{
					"Error 106:\rIncorrect login ( Empty )",
				}
				log.Println(ErrUrl[0])
				fmt.Fprintf(w, ErrUrl[0])
			}else{
				if password == "" {
					ErrUrl := [...]string{
						"Error 105:\rIncorrect password ( Empty )",
					}
					log.Println(ErrUrl[0])
					fmt.Fprintf(w, ErrUrl[0])
				}else{
					//rest api auth

					resp, err := http.Get("https://vavylon1cardinal.herokuapp.com/api?title=auth&token=cardinal&login=login&password=password")
					if err != nil {
						print(err)
					}
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						print(err)
					}
					//fmt.Print(string(body))
					var res map[string]interface{}
					json.Unmarshal([]byte(body), &res)
					if res["Status"] == "OK"{
						log.Println("Nice request")
					}

					if res["Status"] == "Err, bad token"{
						OtherErrors := [...]string{
							"Error 202:\rIncorrect Token Cardinal",
						}
						log.Println(OtherErrors[0])
						fmt.Fprintf(w, OtherErrors[0])
					}

					if res["Status"] == "Err, bad login"{
						OtherErrors := [...]string{
							"Error 206:\rIncorrect login",
						}
						log.Println(OtherErrors[0])
						fmt.Fprintf(w, OtherErrors[0])
					}

					if res["Status"] == "Err, bad password"{
						OtherErrors := [...]string{
							"Error 205:\rIncorrect password",
						}
						log.Println(OtherErrors[0])
						fmt.Fprintf(w, OtherErrors[0])
					}

				}
			}
		}
	}

	if title == "reg" {
		if token != "cardinal" {
			OtherErrors := [...]string{
				"Error 202:\rIncorrect Token Cardinal",
			}
			log.Println(OtherErrors[0])
			fmt.Fprintf(w, OtherErrors[0])
		}else{
			//send req cardinal
			if login == "" {
				ErrUrl := [...]string{
					"Error 106:\rIncorrect login ( Empty )",
				}
				log.Println(ErrUrl[0])
				fmt.Fprintf(w, ErrUrl[0])
			}else{
				if password == "" {
					ErrUrl := [...]string{
						"Error 105:\rIncorrect password ( Empty )",
					}
					log.Println(ErrUrl[0])
					fmt.Fprintf(w, ErrUrl[0])
				}else{
					if email == ""{
						ErrUrl := [...]string{
							"Error 103:\rIncorrect email ( Empty )",
						}
						log.Println(ErrUrl[0])
						fmt.Fprintf(w, ErrUrl[0])
					}else{
						//rest api auth

						resp, err := http.Get("https://vavylon1cardinal.herokuapp.com/api?title=auth&token=cardinal&login=login&password=password")
						if err != nil {
							print(err)
						}
						defer resp.Body.Close()
						body, err := ioutil.ReadAll(resp.Body)
						if err != nil {
							print(err)
						}
						//fmt.Print(string(body))
						var res map[string]interface{}
						json.Unmarshal([]byte(body), &res)
						if res["Status"] == "OK"{
							log.Println("Nice request")
						}

						if res["Status"] == "Err, bad email"{
							OtherErrors := [...]string{
								"Error 203:\rIncorrect email",
							}
							log.Println(OtherErrors[0])
							fmt.Fprintf(w, OtherErrors[0])
						}

						if res["Status"] == "Err, bad login"{
							OtherErrors := [...]string{
								"Error 206:\rIncorrect login",
							}
							log.Println(OtherErrors[0])
							fmt.Fprintf(w, OtherErrors[0])
						}

						if res["Status"] == "Err, bad password"{
							OtherErrors := [...]string{
								"Error 205:\rIncorrect password",
							}
							log.Println(OtherErrors[0])
							fmt.Fprintf(w, OtherErrors[0])
						}

					}
				}
			}
		}
	}

}
