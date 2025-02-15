package routes_user

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

	app.Router.Handle("/user/create", middlewares.Chain(
		http.HandlerFunc(userRoutes.CreateRoute),

		middlewares.CorsMiddleware("GET"),
	))
}
