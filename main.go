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

	for i := 0; i < 100; i++ {
		user := models.User{
			Name: fmt.Sprintf("User %d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
		}

		err = userModel.Insert(&user)
		if err != nil {
			log.Fatalln(err)
		}
	}

	filter := models.Filter{
		PageSize: 2,
		Page: 2,
	}
	users, _, err := userModel.All(filter)
	if err != nil {
		log.Fatalln(err)
	}

	for _, user := range users {
		fmt.Printf("%d: %s %s\n", user.ID, user.Name, user.Email)
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