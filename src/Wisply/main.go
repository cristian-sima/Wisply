package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func lala(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Da ma merge!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/mauritus", lala)
	http.ListenAndServe(":8000", nil)
}
