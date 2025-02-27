package routes_api_prod

import (
	"net/http"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/app"
	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/middlewares"
)

type RoutesProd struct {
	App *app.Application
}

func RegisterRoutes(app *app.Application) {
	prodRoutes := &RoutesProd{App: app}

	app.Router.Handle("/test", middlewares.Chain(
		http.HandlerFunc(prodRoutes.CreateRoute),

		middlewares.CorsMiddleware("GET"),
	))
}
