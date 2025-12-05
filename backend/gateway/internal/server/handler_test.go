package server_test

import (
	"context"
	"encoding/json"
	"gateway/internal/mocks"
	"gateway/internal/server"
	"gateway/internal/services/users"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := mocks.NewMockUserService(ctrl)
	userService.EXPECT().
		Create(context.Background(), *users.NewUser("test1", "test1@gmail.com", "test1-pass")).
		Return(int64(1), nil)

	body := `{"name":"test1","email":"test1@gmail.com","password":"test1-pass"}`
	req := httptest.NewRequest(http.MethodPost, "/users/create", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler, err := server.NewHttpHandler(userService)
	require.NoError(t, err)
	handler.ServeHTTP(rr, req)

	require.Equal(t, rr.Code, http.StatusCreated)
	require.Equal(t, `{"id":1}`, strings.TrimSpace(rr.Body.String()))
}

func TestFindUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	email := "test1@gmail.com"
	expectedUser := users.NewUser("test1", email, "test1-pass")
	userService := mocks.NewMockUserService(ctrl)
	userService.EXPECT().
		Find(context.Background(), email).
		Return(expectedUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/users?email="+email, nil)
	rr := httptest.NewRecorder()

	handler, err := server.NewHttpHandler(userService)
	require.NoError(t, err)
	handler.ServeHTTP(rr, req)

	require.Equal(t, rr.Code, http.StatusOK)

	var actualUser users.User
	err = json.NewDecoder(rr.Body).Decode(&actualUser)
	require.NoError(t, err)

	require.Equal(t, *expectedUser, actualUser)
}
