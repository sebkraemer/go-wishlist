package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sebkraemer/go-wishlist/pkg/config"
	"github.com/sebkraemer/go-wishlist/pkg/handlers"
	"github.com/sebkraemer/go-wishlist/pkg/render"
)

const portNumber = ":8080"

// wishlistItem internal resource
type wishlistItem struct {
	// what about _id handling?
	// ID          primitive.ObjectID `bson:"_id, omitempty"` // will create 000000 ids if not set
	Owner       string    `bson:"owner"`
	Description string    `bson:"desc"`
	Date        time.Time `bson:"date"`
}

// Wish Resource type for REST API
type Wish struct {
	Description string `json:"desc"  bson:"desc"`
}

func main() {
	var app config.AppConfig
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
