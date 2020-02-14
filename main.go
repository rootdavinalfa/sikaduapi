/*
 * Copyright (c) 2019. dvnlabs.ml
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * API For sikadu.unbaja.ac.id
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"unbajaUAPI/core/Student"
	jwt "unbajaUAPI/libs"
	"unbajaUAPI/model"
)

const USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"

type res struct {
	Error       bool
	Message     string
	Description string
	SourceCode  string
	Creator     string
}
type data struct {
	Error   bool
	Message string
	Data    interface{}
}

func main() {
	initRoute()
}
func initRoute() {
	var port = os.Getenv("PORT")
	if port == "" {
		println("Using default port 8080")
		port = "8080"
	}
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/login/mahasiswa", loginHandlerMhs).Methods("POST")
	router.HandleFunc("/login/dosen", loginHandlerDsn)
	router.HandleFunc("/mahasiswa/info/{token}", mhsInfoHandler).Methods("GET")
	router.HandleFunc("/mahasiswa/schedule/{token}", mhsScheduleList).Methods("GET")
	//Schedule quart is 1/2 1 for odd semester 2 for even semester, Year is academic year
	router.HandleFunc("/mahasiswa/schedule/{year}/{quart}/{token}", mhsSchedule).Methods("GET")
	router.HandleFunc("/mahasiswa/grade/{year}/{quart}/{token}", mhsGradeDetail).Methods("GET")
	//Handler for get summary grade
	router.HandleFunc("/mahasiswa/grade/summary/{token}", mhsGradeSummary).Methods("GET")
	//Handler for get finance status
	router.HandleFunc("/mahasiswa/finance/{token}", mhsFinance).Methods("GET")

	// Handle all preflight request
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	router.NotFoundHandler = http.HandlerFunc(notFound)
	router.StrictSlash(true)
	color.Warn.Println("Connected to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	home := res{
		Error:   true,
		Message: "Not found matched endpoint",
	}
	var homeJson = string(MustMarshal(home))
	_, _ = fmt.Fprint(w, homeJson)
}

func mhsScheduleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	param := mux.Vars(r)
	token := param["token"]
	data := Student.GetStudentScheduleListHub(w, token)
	_, _ = fmt.Fprint(w, string(MustMarshal(data)))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	home := res{
		Error:       false,
		Message:     "(c)2019 dvnlabs.ml -> Davin Alfarizky Putra Basudewa",
		Description: "This is Unofficial API for sikadu.unbaja.ac.id JSON Ready! 🎏 🈷 . You must /login/mahasiswa before any request and save token.The token is valid for 1 hour.Because sikadu not even have error handler,we cannot determine token is valid or not",
		Creator:     "https://dvnlabs.ml",
		SourceCode:  "https://github.com/rootdavinalfa/sikaduapi",
	}
	var homeJson = string(MustMarshal(home))
	_, _ = fmt.Fprint(w, homeJson)
}
func loginHandlerMhs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	if err := r.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	ErrS := true
	Mesage := ""
	Descc := ""
	/*user is numeric 10digit your NIM*/
	user := r.FormValue("user")
	password := r.FormValue("password")
	if user != "" && password != "" {
		formLogin := url.Values{
			"user":     {user},
			"password": {password},
			"as":       {"Mahasiswa"},
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := http.PostForm("http://sikadu.unbaja.ac.id/main/login", formLogin)
		if err != nil {
			println(err.Error())
		}
		if resp != nil {
			for _, cookie := range resp.Cookies() {
				if cookie.Name == "ci_session" {
					token := model.LoginAuth{
						User:   user,
						Cookie: cookie.Value,
					}
					isError, data := jwt.NewToken(token)
					if isError {
						Mesage = "JWT ERROR"
					} else {
						ErrS = false
						Descc = data
					}
				}
			}
		}
	} else if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		Mesage = "user blank/empty"
	} else if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		Mesage = "password blank/empty"
	}
	type tokenizer struct {
		Error   bool
		Message string
		Token   string
	}
	response := tokenizer{
		Error:   ErrS,
		Message: Mesage,
		Token:   Descc,
	}
	_, _ = fmt.Fprint(w, string(MustMarshal(response)))
}

func loginHandlerDsn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	home := res{
		Error:       false,
		Message:     "This endpoint not implemented caused by not enough information",
		Description: "🆗 This is sikadu.unbaja.ac.id Unofficial API written by Davin Alfarizky Putra Basudewa - 1101171082. (c)2019 dvnlabs.ml",
	}
	var homeJson = string(MustMarshal(home))
	_, _ = fmt.Fprint(w, homeJson)

}
func mhsInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	params := mux.Vars(r)
	token := params["token"]
	data := Student.GetStudentInfoHub(w, token)

	_, _ = fmt.Fprint(w, string(MustMarshal(data)))
}
func mhsSchedule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	params := mux.Vars(r)
	token := params["token"]
	year := params["year"]
	quart := params["quart"]
	data := Student.GetStudentScheduleHub(w, token, year, quart)
	_, _ = fmt.Fprint(w, string(MustMarshal(data)))
}

func mhsGradeDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	params := mux.Vars(r)
	token := params["token"]
	year := params["year"]
	quart := params["quart"]
	data := Student.GetStudentGradeDetailHub(w, token, year, quart)
	_, _ = fmt.Fprint(w, string(MustMarshal(data)))
}
func mhsGradeSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	params := mux.Vars(r)
	token := params["token"]
	data := Student.GetStudentGradeSummaryHub(w, token)
	_, _ = fmt.Fprint(w, string(MustMarshal(data)))
}

func mhsFinance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	LogConsoleHttpReq(r)
	params := mux.Vars(r)
	token := params["token"]
	data := Student.GetStudentFinanceStatus(w, token)
	_, _ = fmt.Fprint(w, string(MustMarshal(data)))
}

func MustMarshal(data interface{}) []byte {
	out, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return out
}
func LogConsoleHttpReq(r *http.Request) {
	color.Cyan.Println(r.Method + " : " + r.Proto + " [" + r.Host + r.URL.String() + "] Requested by: " + r.RemoteAddr + " At:->" + time.Now().String())
}
