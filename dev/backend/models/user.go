package models

import "errors"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// errors

var ErrorInvalidUser = errors.New("Invalid user")
