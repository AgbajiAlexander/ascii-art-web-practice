package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", name)

}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == http.MethodPost {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read Body", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Character count: %d", len(bodyBytes))
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	aNum, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "Invalid number for 'a'", http.StatusBadRequest)
		return
	}

	bNum, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "Invalid number for 'b'", http.StatusBadRequest)
		return
	}

	var result int
	switch op {
	case "add":
		result = aNum + bNum
	case "subtract":
		result = aNum - bNum
	case "multiply":
		result = aNum * bNum
	default:
		http.Error(w, "Unknown operation", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %d", result)

}

func agentHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	if userAgent == "" {
		userAgent = "Unknown"
	}
	fmt.Fprintf(w, "You are visiting us using: %s", userAgent)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	clientKey := r.Header.Get("X-API-Key")
	if clientKey != "secret123" {
		http.Error(w, "Unauthorized:", http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, "Welcome to the secure DASHBOARD!")
}

func legacyHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func v2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}
func methodInspectorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yon made a %s request", r.Method)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/agent", agentHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/legacy", legacyHandler)
	http.HandleFunc("/v2", v2Handler)
	http.HandleFunc("/method-inspector", methodInspectorHandler)

	fmt.Print("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
