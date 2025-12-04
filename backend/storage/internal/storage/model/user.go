package model

import "time"

type User struct {
	Name     string
	Email    string
	Created  time.Time
	Password string
}

func NewUser(name string, email string, password string) *User {
	return &User{Name: name, Email: email, Created: time.Now(), Password: password}
}
