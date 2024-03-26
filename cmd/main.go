package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/jackcollins434/guess-the-spotify-song/cmd/view"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", templ.Handler(view.Home()).ServeHTTP)
	http.ListenAndServe(":3001", r)
}
