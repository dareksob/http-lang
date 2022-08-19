package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Msg struct {
	Msg string
}

func accessLog(r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
}

func prinftJson(w http.ResponseWriter, m Msg) {
	jsonStatus, err := json.Marshal(m)

	if (err != nil) {
			fmt.Fprintf(w, "Error", err)
	} else {
			fmt.Fprintf(w, string(jsonStatus))
	}
}


func main() {
	
	// GET /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		accessLog(r)
		fmt.Fprintf(w, "hello world")
	})

	// GET /get-json
	http.HandleFunc("/get-json", func(w http.ResponseWriter, r *http.Request) {
		accessLog(r)
		prinftJson(w, Msg{"hello world"})
	})

	// GET /get/:id
	// @todo use an parameter
	http.HandleFunc("/get/name", func(w http.ResponseWriter, r *http.Request) {
		accessLog(r)
		prinftJson(w, Msg{"hello name"})
	})

	log.Printf("Go core webserver running. Access it at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}