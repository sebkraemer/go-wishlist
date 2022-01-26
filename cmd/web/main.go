package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sebkraemer/go-wishlist/pkg/config"
	"github.com/sebkraemer/go-wishlist/pkg/handlers"
	"github.com/sebkraemer/go-wishlist/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	// todo change when in production/ make runtime configurable
	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // should be true in production for tls

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	// apply repository pattern to set config
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	fmt.Printf("Starting application on port %s\n", portNumber)
	log.Fatal(srv.ListenAndServe())
}
