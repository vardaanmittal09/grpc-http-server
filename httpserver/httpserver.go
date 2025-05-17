package httpserver

import (
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Fprint(w, string(runes))
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)
	if errA != nil || errB != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%d", a+b)
}

func StartHTTPServer(addr string) error {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/reverse", reverseHandler)
	http.HandleFunc("/sum", sumHandler)
	return http.ListenAndServe(addr, nil)
}
