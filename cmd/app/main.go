package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wyle09/go-bookings-web-app/pkg/config"
	"github.com/wyle09/go-bookings-web-app/pkg/handler"
	"github.com/wyle09/go-bookings-web-app/pkg/render"
)

const portNumber = ":8080"

// Main application function
func main() {

	app := config.AppConfig{}

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache:", err)
	}

	app.TemplateCache = templateCache

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	// Pass the address location of the app.
	render.NewTemplates(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
