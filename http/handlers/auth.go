package handlers

import (
	"katinka/config"
	"katinka/services/container"
	"net/http"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		internalError(err, w)
		return
	}

	dbName := r.FormValue("collection")
	usr := r.FormValue("user")
	pwd := r.FormValue("password")

	ok := container.Connect(usr, pwd, dbName)

	if ok {
		http.Redirect(w, r, config.BasePath, http.StatusFound)
	} else {
		http.Redirect(w, r, config.AuthPath, http.StatusFound)
	}
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	container.Disconnect()
	http.Redirect(w, r, config.AuthPath, http.StatusFound)

	return
}
