package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kjuchniewicz/go-web/pkg/config"
	"github.com/kjuchniewicz/go-web/pkg/handlers"
	"github.com/kjuchniewicz/go-web/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Startujemy na porcie %s\n", portNumber[1:])

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
