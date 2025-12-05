package main

import (
	"gateway/internal/server"
	"gateway/internal/services/users"

	"log"
	"net/http"
)

func main() {
	userSvc, err := users.NewUserServiceClient()
	if err != nil {
		log.Fatal(err)
	}
	defer userSvc.Close()

	handler, err := server.NewHttpHandler(userSvc)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", handler)
}
