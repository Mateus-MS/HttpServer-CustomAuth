package models

import (
	"database/sql"
	"errors"
	"strings"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`

	SessionToken sql.NullString `json:"session_token"`
	CSRFToken    sql.NullString `json:"csrf_token"`
}

func (user *User) Validate() error {
	if user.Username == "" {
		return ErrorInvalidUser
	}

	if strings.Contains(user.Username, " ") {
		return ErrorInvalidUser
	}

	if user.PasswordHash == "" {
		return ErrorInvalidUser
	}

	if strings.Contains(user.PasswordHash, " ") {
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
		PasswordHash: user.PasswordHash,
		SessionToken: user.SessionToken,
		CSRFToken:    user.CSRFToken,
	}
}

// errors
var (
	ErrorInvalidUser       = errors.New("Invalid user")
	ErrorUserNotFound      = errors.New("User not found")
	ErrorUserAlreadyExists = errors.New("User already exists")
)
