package routes_pages

import (
	"net/http"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/app"
	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/middlewares"
)

type RoutesPages struct {
	App *app.Application
}

func RegisterRoutes(app *app.Application) {
	pagesRoutes := &RoutesPages{App: app}

	app.Router.Handle("/protected", middlewares.Chain(
		http.HandlerFunc(pagesRoutes.ProtectedRoute),

		middlewares.CorsMiddleware("GET"),
		middlewares.Authenticate(app.DB),
	))
}
