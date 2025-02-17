package app

import (
	"database/sql"
	"net/http"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services"
	service_prod "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/prod"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
)

type Application struct {
	DB     *sql.DB
	Router *http.ServeMux

	Services map[string]services.Service
}

func NewApplication() *Application {
	app := Application{
		DB:     GetInstance(),
		Router: http.NewServeMux(),
	}

	app.Services = make(map[string]services.Service)

	app.Services["user"] = &service_user.ServiceUser{DB: app.DB}
	app.Services["prod"] = &service_prod.ServiceProd{DB: app.DB}

	return &app
}
