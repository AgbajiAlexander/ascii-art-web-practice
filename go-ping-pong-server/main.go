package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	fmt.Print("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
