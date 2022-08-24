package main

import (
	"fmt"
	"log"
	"regexp"
	"net/http"
	"encoding/json"
)

type Msg struct {
	Msg string `json:"msg"`
}

func accessLog(r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
}

func prinftJson(w http.ResponseWriter, m Msg) {
	jsonStatus, err := json.Marshal(m)

	if (err != nil) {
		fmt.Fprintf(w, "Error", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonStatus))
	}
}

func main() {
	// GET /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		accessLog(r)

		if (r.URL.Path == "/") {
			fmt.Fprintf(w, "hello world")
		} else {
			http.NotFound(w, r)
		}
	})

	// GET /get-json
	http.HandleFunc("/get-json", func(w http.ResponseWriter, r *http.Request) {
		accessLog(r)
		prinftJson(w, Msg{"hello world"})
	})

	// matches for routes
	var getId = regexp.MustCompile("^/get/(.*)")
	

	// GET /get-json
	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		accessLog(r)

		if (getId.MatchString(r.URL.Path)) {
			match := getId.FindStringSubmatch(r.URL.Path)
			prinftJson(w, Msg{"hello " + match[1]})
		} else {
			http.NotFound(w, r)
		}
	})

	log.Printf("Go core webserver running. Access it at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}