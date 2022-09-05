package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jamessom/go-web-basics/models"

	_ "github.com/lib/pq"
)

type application struct {
	Models models.Models
}

func main() {

	dns := "postgres://gopguser:gopgpass@go-db.gowebbasics.com/gopgdb?sslmode=disable"

	db, err := connectToDb(dns)
	if err != nil {
		log.Fatalln(err)
	}

	app := application{
		Models: models.NewModel(db),
	}

	fmt.Println("Starting application")

	err = app.server()
	if err != nil {
		log.Fatalln(err)
	}
}

func connectToDb(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}