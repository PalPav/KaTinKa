package middlewares

import (
	"katinka/config"
	"katinka/services/container"
	"net/http"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isConn := container.DB() != nil

		// skip auth for static files
		if len(r.URL.Path) > 8 && r.URL.Path[:8] == "/static/" {
			h.ServeHTTP(w, r)
			return
		}

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
