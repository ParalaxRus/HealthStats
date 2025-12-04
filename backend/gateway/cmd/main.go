package main

import (
	"gateway/internal/server"
	"gateway/internal/services/users"

	"log"
	"net/http"
)

func main() {
	usersClient, err := users.NewUsersClient()
	if err != nil {
		log.Fatal(err)
	}
	defer usersClient.Close()

	handler, err := server.NewHttpHandler(usersClient)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", handler)
}
