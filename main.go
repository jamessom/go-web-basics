package main

import (
	"database/sql"
	"fmt"
	"go-with-docker-compose/models"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	dns := "postgres://gopguser:gopgpass@go-db.gowithdockercompose.com/gopgdb?sslmode=disable"

	db, err := connectToDb(dns)
	if err != nil {
		log.Fatalln(err)
	}

	userModel := models.UsersModel{
		DB: db,
	}
	user := models.User{
		Name: "James",
		Email: "james@a.com",
	}

	err = userModel.Insert(&user)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Created user with ID %d\n", user.ID)
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