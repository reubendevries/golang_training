package main

import (
	"fmt"
	"golang_training/Section-02/05-BasicStyling/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function.
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
