package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"rest_imageboard/internal/config"
	"rest_imageboard/internal/server/handlers/threads"
	"rest_imageboard/internal/storage/query"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	cfg := config.GetConfig()
	log.Println("succesfully loaded config")

	pgInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)

	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		log.Fatal("couldn't open database: " + err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("couldn't ping the database: " + err.Error())
	}

	log.Println("succesfully connected to postgresql")
	log.Print(query.CreateTablesIfNotCreated(db))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/threads", func(r chi.Router) {
		r.Get("/", threads.Get(db))
		r.Post("/", threads.Save(db))
		r.Delete("/{name}", threads.Delete(db))
	})
	http.ListenAndServe(":8081", r)
}
