package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// GET /
	mux.HandleFunc("/", w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello world")
}


	http.ListenAndServe(":8080", mux)
}