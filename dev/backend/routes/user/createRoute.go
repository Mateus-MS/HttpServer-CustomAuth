package routes_user

import (
	"net/http"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
)

func (app *RoutesUser) CreateRoute(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := r.FormValue("username")
	email := r.FormValue("email")
	pass := r.FormValue("password")

	registerOBJ := models.User{
		Username: user,
		Email:    email,
		Password: pass,
	}

	if err := service_user.Create(&registerOBJ, app.App.DB); err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "tb_user_username_key"` {
			w.WriteHeader(http.StatusConflict)
			return
		}

		println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
