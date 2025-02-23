package routes_api_user

import (
	"net/http"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/app"
	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/middlewares"
)

type RoutesUser struct {
	App *app.Application
}

func RegisterRoutes(app *app.Application) {
	userRoutes := &RoutesUser{App: app}

	app.Router.Handle("/api/user/register", middlewares.Chain(
		http.HandlerFunc(userRoutes.Create),

		middlewares.CorsMiddleware("POST"),
	))

	app.Router.Handle("/api/user/login", middlewares.Chain(
		http.HandlerFunc(userRoutes.Login),

		middlewares.CorsMiddleware("POST"),
	))
}
