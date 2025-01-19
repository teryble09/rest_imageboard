package main

import (
	"database/sql"
	"fmt"
	"log"

	"rest_imageboard/internal/config"

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
		log.Fatal("couldn't connect to a database: " + err.Error())
	}
	defer db.Close()
	log.Println("succesfully connected to postgresql")



}