package routes_api_user

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/utils"
)

func (app *RoutesUser) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// It can be the username or the email
	identifyier := r.FormValue("identifyer")
	password := r.FormValue("password")

	// Create the search user object
	loginOBJ := models.User{}
	if strings.Contains(identifyier, "@") {
		loginOBJ.Email = identifyier
	} else {
		loginOBJ.Username = identifyier
	}

	// Query the user
	userOBJ, err := service_user.Search(&loginOBJ, app.App.DB)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If password received miss match the one in DB
	if !utils.CheckPassordHash(password, userOBJ.PasswordHash) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	err = service_user.Update(userOBJ, &models.User{
		SessionToken: sql.NullString{Valid: true, String: sessionToken},
		CSRFToken:    sql.NullString{Valid: true, String: csrfToken},
	}, app.App.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		Expires:  time.Now().Add(time.Hour * 24),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		Expires:  time.Now().Add(time.Hour * 24),
	})
}
