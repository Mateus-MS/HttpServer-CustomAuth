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
	app := Application{
		DB:     GetInstance(),
		Router: http.NewServeMux(),
	}

	return &app
}
