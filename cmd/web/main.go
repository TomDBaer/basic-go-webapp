package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TomDBaer/basic-go-webapp/pkg/config"
	"github.com/TomDBaer/basic-go-webapp/pkg/handlers"
	"github.com/TomDBaer/basic-go-webapp/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	// Hier wird die config geladen ("pkg/config")
	var app config.AppConfig

	// tt := render.CreateTemplateCache()

	templateCacheAd, err := render.CreateTemplateCacheAdvanced()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCacheAdvanced = templateCacheAd

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/divide", handlers.Divide)

	fmt.Printf("Starting application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
