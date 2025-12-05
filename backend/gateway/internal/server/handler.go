package server

import (
	"encoding/json"
	"errors"
	"gateway/internal/services/users"
	"net/http"
	"strings"
)

type handler struct {
	svc users.UserService
}

func NewHttpHandler(svc users.UserService) (http.Handler, error) {
	if svc == nil {
		return nil, errors.New("user service client is not set")
	}
	return &handler{svc: svc}, nil
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if isCreateUserRequest(req) {
		h.createUserHandler(w, req)
		return
	}

	if isFindUserRequest(req) {
		h.findUserHandler(w, req)
		return
	}

	if isListUsersRequest(req) {
		h.svc.List(req.Context(), "")
		return
	}

	http.NotFound(w, req)
}

func isCreateUserRequest(req *http.Request) bool {
	return req.Method == "POST" && strings.HasPrefix(req.URL.Path, "/users")
}

func (h *handler) createUserHandler(w http.ResponseWriter, req *http.Request) {
	user, err := getUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.svc.Create(req.Context(), *user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]any{
		"id": id,
	})
}

func (h *handler) findUserHandler(w http.ResponseWriter, req *http.Request) {
	email, err := getEmail(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.svc.Find(req.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func isFindUserRequest(req *http.Request) bool {
	return (req.Method == "GET") &&
		strings.HasPrefix(req.URL.Path, "/users") &&
		(len(req.URL.Query().Get("email")) > 0)
}

func isListUsersRequest(req *http.Request) bool {
	return req.Method == "GET" && strings.HasPrefix(req.URL.Path, "/users")
}

func getUser(req *http.Request) (*users.User, error) {
	var u users.User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func getEmail(req *http.Request) (string, error) {
	return req.URL.Query().Get("email"), nil
}
