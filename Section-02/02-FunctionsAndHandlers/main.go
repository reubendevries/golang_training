package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is the home page")
	if err != nil {
		fmt.Errorf("Error %v", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is the about page")
	if err != nil {
		fmt.Errorf("Error %v", err)
	}
}

func AddValues(x, y int) (int, error) {
	return x + y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	_ = http.ListenAndServe(":8080", nil)
}
