package routes_pages

import "net/http"

func (app *RoutesPages) ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Protected Page</title>
	</head>
	<body>
		<h1>Welcome to the Protected Page</h1>
		<p>This content is only available to authorized users.</p>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
