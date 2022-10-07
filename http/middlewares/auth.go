package middlewares

import (
	"katinka/config"
	"katinka/services/container"
	"net/http"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isConn := container.DB() != nil

		if isConn && r.URL.Path == config.AuthPath {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		if !isConn && r.URL.Path != config.AuthPath {
			http.Redirect(w, r, config.AuthPath, http.StatusTemporaryRedirect)
			return
		}

		h.ServeHTTP(w, r)
	})
}
