package middlewares

import (
	"database/sql"
	"net/http"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
)

// Authenticate is a middleware that checks if the user is authenticated
// It will stop the request if the user is not authenticated
// To be authenticated the user must have a session_token & the csrf_token in the cookies
func Authenticate(db *sql.DB) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Query both cookies
			// If any of them is not found, means that the user is not authenticated
			session_cookie, err := r.Cookie("session_token")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			csrf_cookie, err := r.Cookie("csrf_token")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Search for the user with the session_token
			// So if not found, means that the user or is not logged or does not exist
			searchedOBJ, err := service_user.Search(
				&models.User{
					SessionToken: sql.NullString{
						String: session_cookie.Value,
						Valid:  true,
					},
				},
				db)

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// If the csrf_token is not the same as the one in the DB
			// It means that the user is trying to do a CSRF attack
			// So we will stop the request
			if csrf_cookie.Value != searchedOBJ.CSRFToken.String {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
