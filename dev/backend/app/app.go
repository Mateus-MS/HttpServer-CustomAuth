package app

import (
	"database/sql"
	"net/http"
)

type Application struct {
	DB     *sql.DB
	Router *http.ServeMux
}

func NewApplication() *Application {
	return &Application{
		DB:     GetInstance(),
		Router: http.NewServeMux(),
	}
}
