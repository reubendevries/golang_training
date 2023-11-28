package main

import (
	"errors"
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

// Divide is the divide page handler.
func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 10.0
	f, err := divideValues(x, y)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Can't divide by zero.")
		return
	}
	_, _ = fmt.Fprintf(w, "%f divided by %f is %f", x, y, f)
}

// addValues adds two integers and returns the sum.
func addValues(x, y int) int {
	return x + y
}

// divideValues divides two intergers and returns the sum.
func divideValues(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0, errors.New("Can't divide by zero.")
	} else {
		return x / y, nil
	}
}

// main is the main application function.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Printf("Starting application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
