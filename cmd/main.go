package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/jackcollins434/guess-the-spotify-song/internal/model"
	"github.com/jackcollins434/guess-the-spotify-song/internal/view"
	_ "github.com/joho/godotenv/autoload"
	"github.com/michaeljs1990/sqlitestore"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DB_PATH = "internal/db/app.db"

var store *sqlitestore.SqliteStore

func init() {
	var err error
	store, err = sqlitestore.NewSqliteStore(DB_PATH, "sessions", "/", 3600, []byte(os.Getenv("APP_SECRET")))
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	r := chi.NewRouter()
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "foobar")
		session.Values["bar"] = "baz"
		err = session.Save(r, w)
		fmt.Println("err", err)

		templ.Handler(view.Home()).ServeHTTP(w, r)
	})

	http.ListenAndServe(":3001", r)
}
