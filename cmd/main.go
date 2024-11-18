package main

import (
	"database/sql"
	"debtsapp/internal/env"
	"debtsapp/internal/storage"
	_ "github.com/lib/pq"
	"log"
)

type Configuration struct {
	Port string
}

type Application struct {
	Storage       *storage.Storage
	Configuration Configuration
}

func (a *Application) Start() {

	StartWebserver(a)
}

func main() {

	db, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	app := &Application{
		Storage: storage.NewStorage(db),
		Configuration: Configuration{
			Port: env.GetString("PORT", "8080"),
		},
	}

	app.Start()

}
