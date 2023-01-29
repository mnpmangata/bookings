package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mnpmangata/bookings/pkg/config"
	"github.com/mnpmangata/bookings/pkg/handlers"
)

func routes(app *config.Appconfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole) -deleted
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
