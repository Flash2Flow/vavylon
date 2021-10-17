package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseTest struct {
	Status		string
	Code string
}

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
	//UserKeyByte := query.Get("key")
	PasswordByte := query.Get("password")
	LoginByte := query.Get("login")


	title := string(TitleByte[:])		//err id 1
	token := string(TokenByte[:])		//err id 2
	email := string(EmailByte[:])		//err id 3
	//userkey := string(UserKeyByte[:])	//err id 4
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

/*
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
					if login == "" {
						ErrUrl := [...]string{
							"Error 106:\rIncorrect login ( Empty )",
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


					}
				}
				}
			}
		}
 */

	if title == "auth" {
		if token != "cardinal" {
			OtherErrors := [...]string{
				"Error 202:\rIncorrect Token Cardinal",
			}
			log.Println(OtherErrors[0])
			fmt.Fprintf(w, OtherErrors[0])

			responsetest := ResponseTest{"ERR",  "202 Incorrect Token Cardinal"}

			js, err := json.Marshal(responsetest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			log.Println(js)
		}else{
			//send req cardinal
			if login == "" {
				ErrUrl := [...]string{
					"Error 106:\rIncorrect login ( Empty )",
				}
				log.Println(ErrUrl[0])
				fmt.Fprintf(w, ErrUrl[0])

				responsetest := ResponseTest{"ERR",  "??? Empty url login"} //change err code

				js, err := json.Marshal(responsetest)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
				log.Println(js)
			}else{
				if password == "" {
					ErrUrl := [...]string{
						"Error 105:\rIncorrect password ( Empty )",
					}
					log.Println(ErrUrl[0])
					fmt.Fprintf(w, ErrUrl[0])

					responsetest := ResponseTest{"ERR",  "??? Empty url password"} //change err code

					js, err := json.Marshal(responsetest)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Header().Set("Content-Type", "application/json")
					w.Write(js)
					log.Println(js)
				}else{
					//rest api auth
					querry := fmt.Sprintf("https://vavylon1cardinal.herokuapp.com/api?title=auth&token=cardinal&login=%s&password=%s", login, password)
					resp, err := http.Get(querry)
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

					if res["Status"] == "ERR"{
						OtherErrors := [...]string{
							"ERRRRRRRRRRRRRRR",
						}
						log.Println(OtherErrors[0])
						fmt.Fprintf(w, OtherErrors[0])

						if res["Code"] == "201 Incorrect Token Cardinal"{
							responsetest := ResponseTest{"ERR",  "201 Incorrect Token Cardinal"}

							js, err := json.Marshal(responsetest)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}

							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
							log.Println(js)
						}

						if res["Code"] == "??? Empty url login"{
							responsetest := ResponseTest{"ERR",  "??? Empty url login"}

							js, err := json.Marshal(responsetest)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}

							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
							log.Println(js)
						}
						if res["Code"] == "??? Empty url password"{
							responsetest := ResponseTest{"ERR",  "??? Empty url password"}

							js, err := json.Marshal(responsetest)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}

							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
							log.Println(js)
						}

						if res["Code"] == "304 User not found"{
							responsetest := ResponseTest{"ERR",  "304 User not found"}

							js, err := json.Marshal(responsetest)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}

							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
							log.Println(js)
						}

						if res["Code"] == "302 Bad password"{
							responsetest := ResponseTest{"ERR",  "302 Bad password"}

							js, err := json.Marshal(responsetest)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}

							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
							log.Println(js)
						}

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

			responsetest := ResponseTest{"ERR",  "291 Bad token cardinal"}

			js, err := json.Marshal(responsetest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			log.Println(js)
		}else{
			//send req cardinal
			if login == "" {
				ErrUrl := [...]string{
					"Error 106:\rIncorrect login ( Empty )",
				}
				log.Println(ErrUrl[0])
				fmt.Fprintf(w, ErrUrl[0])
				responsetest := ResponseTest{"ERR",  "Incorrect login ( Empty )"}

				js, err := json.Marshal(responsetest)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
				log.Println(js)
			}else{
				if password == "" {
					ErrUrl := [...]string{
						"Error 105:\rIncorrect password ( Empty )",
					}
					log.Println(ErrUrl[0])
					fmt.Fprintf(w, ErrUrl[0])
					responsetest := ResponseTest{"ERR",  "Incorrect password ( Empty )"}

					js, err := json.Marshal(responsetest)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Header().Set("Content-Type", "application/json")
					w.Write(js)
					log.Println(js)
				}else{
					if email == ""{
						ErrUrl := [...]string{
							"Error 103:\rIncorrect email ( Empty )",
						}
						log.Println(ErrUrl[0])
						fmt.Fprintf(w, ErrUrl[0])
						responsetest := ResponseTest{"ERR",  "Incorrect email ( Empty )"}

						js, err := json.Marshal(responsetest)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}

						w.Header().Set("Content-Type", "application/json")
						w.Write(js)
						log.Println(js)
					}else{
						//rest api reg
						querry := fmt.Sprintf("https://vavylon1cardinal.herokuapp.com/api?title=reg&token=cardinal&login=%s&password=%s&email=%s", login, password, email)
						resp, err := http.Get(querry)
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
							responsetest := ResponseTest{"OK",  "........."}

							js, err := json.Marshal(responsetest)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}

							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
							log.Println(js)
						}

						if res["Code"] == "999 Unknown error ( бтв такой пользователь уже есть )"{


							responsetest := ResponseTest{"ERR",  "999 Unknown error ( бтв такой пользователь уже есть )"}

							js, err := json.Marshal(responsetest)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}

							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
							log.Println(js)
						}

						if res["Code"] == "999 Unknown error ( бтв такой пользователь уже есть )"{

							responsetest := ResponseTest{"ERR",  "999 Unknown error ( бтв такой пользователь уже есть )"}

							js, err := json.Marshal(responsetest)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}

							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
							log.Println(js)
						}

					}
				}
			}
		}
	}

	//ds integration



}
