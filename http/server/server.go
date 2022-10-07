package server

import (
	"github.com/gorilla/mux"
	"katinka/http/middlewares"
	"katinka/http/routes"
	"net/http"
)

func Run() {
	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)
	routes.RegisterWeb(router)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//graceful shutdown ?
	err := http.ListenAndServe(":80", router)
	if err != nil {
		panic(err)
	}
}
