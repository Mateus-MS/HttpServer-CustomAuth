package models

import (
	"database/sql"
	"errors"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`

	SessionToken sql.NullString `json:"session_token"`
}

func (user *User) Validate() error {
	if user.Username == "" {
		return ErrorInvalidUser
	}

	if strings.Contains(user.Username, " ") {
		return ErrorInvalidUser
	}

	if user.Password == "" {
		return ErrorInvalidUser
	}

	if strings.Contains(user.Password, " ") {
		return ErrorInvalidUser
	}

	if user.Email == "" {
		return ErrorInvalidUser
	}

	if strings.Contains(user.Email, " ") {
		return ErrorInvalidUser
	}

	return nil
}

func (user *User) Copy() User {
	return User{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Password:     user.Password,
		SessionToken: user.SessionToken,
	}
}

// errors

var (
	ErrorInvalidUser       = errors.New("Invalid user")
	ErrorUserNotFound      = errors.New("User not found")
	ErrorUserAlreadyExists = errors.New("User already exists")
)
