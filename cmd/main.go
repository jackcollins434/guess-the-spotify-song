package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/jackcollins434/guess-the-spotify-song/internal/model"
	"github.com/jackcollins434/guess-the-spotify-song/internal/view"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("internal/db/app.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	r := chi.NewRouter()
	r.Get("/", templ.Handler(view.Home()).ServeHTTP)
	http.ListenAndServe(":3001", r)
}
