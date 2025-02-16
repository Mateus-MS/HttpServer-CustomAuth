package app

import (
	"database/sql"
	"net/http"

	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
)

type Application struct {
	DB     *sql.DB
	Router *http.ServeMux

	UserService *service_user.ServiceUser
}

func NewApplication() *Application {
	app := Application{
		DB:     GetInstance(),
		Router: http.NewServeMux(),
	}

	app.UserService = &service_user.ServiceUser{DB: app.DB}

	return &app
}
