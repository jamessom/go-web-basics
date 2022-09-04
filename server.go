package main

import (
	"encoding/json"
	"go-with-docker-compose/models"
	"log"
	"net/http"
)

func (app *application) server() error {
	srv := http.Server{
		Handler: app.handlers(),
		Addr: ":8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (app *application) handlers() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.homePage)
	mux.HandleFunc("/users", app.usersPage)

	return mux
}

func (app *application) homePage(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Server started"))
}

func (app *application) usersPage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	filter := models.Filter{Page: 1, PageSize: 20}
	users, metadata, err := app.Models.Users.All(filter)

	if err != nil {
		log.Fatalln(err)
	}

	response := struct {
		Users []models.User
		Metadata models.Metadata
	}{
		Users: users,
		Metadata: metadata,
	}

	json, err := json.Marshal(response)
	if err != nil {
		log.Fatalln(err)
	}

	writer.Write(json)
}