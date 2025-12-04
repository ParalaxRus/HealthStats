package server

import (
	"errors"
	"gateway/internal/services/users"
	"net/http"
	"strings"
)

type handler struct {
	usersClient *users.UsersClient
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if createUserEndpoint(req) {

	} else if listUsersEndpoint(req) {

	}

	http.NotFound(w, req)
}

func createUserEndpoint(req *http.Request) bool {
	return req.Method == "POST" && strings.HasPrefix(req.URL.Path, "/users/create")
}

func listUsersEndpoint(req *http.Request) bool {
	return req.Method == "GET" && strings.HasPrefix(req.URL.Path, "/users/list")
}

func NewHttpHandler(usersClient *users.UsersClient) (http.Handler, error) {
	if usersClient == nil {
		return nil, errors.New("users client is not set")
	}
	return &handler{usersClient: usersClient}, nil
}
