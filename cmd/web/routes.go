package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sebkraemer/go-wishlist/pkg/config"
	"github.com/sebkraemer/go-wishlist/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// todo nosurf csrf, middleware.go
	mux.Use(handlers.SessionLoadAndSave)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
