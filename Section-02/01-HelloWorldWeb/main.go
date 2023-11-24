package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}
