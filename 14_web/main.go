package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello World </h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is all about me!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server running on :4000")
	http.ListenAndServe(":4000", nil)
}
