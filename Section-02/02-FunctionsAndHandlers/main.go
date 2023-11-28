package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

// add values add two intergers and returns the sum.
func addValues(x, y int) int {
	return x + y
}

// main is the main application function.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
