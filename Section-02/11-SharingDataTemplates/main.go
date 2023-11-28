package main

import (
	"fmt"
	"golang_training/Section-02/11-SharingDataTemplates/config"
	"golang_training/Section-02/11-SharingDataTemplates/handlers"
	"golang_training/Section-02/11-SharingDataTemplates/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function.
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache.")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
