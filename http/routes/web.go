package routes

import (
	"github.com/gorilla/mux"
	"katinka/config"
	"katinka/http/handlers"
)

func RegisterWeb(router *mux.Router) {
	//Auth
	router.HandleFunc(config.AuthPath, handlers.HandleAuthPage).Methods("GET")
	router.HandleFunc(config.AuthPath, handlers.HandleAuth).Methods("POST")
	router.HandleFunc("/logout", handlers.HandleLogout).Methods("POST", "GET")

	//Pages
	router.HandleFunc(config.BasePath, handlers.HandleMainPage).Methods("GET")
	router.HandleFunc("/view/rand", handlers.HandleViewRandPage).Methods("GET")
	//router.HandleFunc("/view/tag", handlers.HandleViewRandPage).Methods("GET")
	router.HandleFunc("/upload", handlers.HandleUploadPage).Methods("GET")
	router.HandleFunc("/search", handlers.HandleSearchPage).Methods("GET")
	router.HandleFunc("/upload", handlers.HandleUploadPost).Methods("POST")

	//Images
	router.HandleFunc("/image/{id:[0-9]+}/edit", handlers.HandleViewImagePage).Methods("GET")
	router.HandleFunc("/image/{id:[0-9]+}", handlers.HandleGetImage).Methods("GET")
	router.HandleFunc("/image/{id:[0-9]+}/hit", handlers.HandleImageHit).Methods("POST")
	router.HandleFunc("/image/search", handlers.HandleImageSearch).Methods("GET")

	//Tags
	router.HandleFunc("/tag/search", handlers.HandleTagSearch).Methods("GET")
	router.HandleFunc("/tag/assign", handlers.HandleTagAssign).Methods("POST")
	router.HandleFunc("/tag/detach", handlers.HandleTagDetach).Methods("POST")
	router.HandleFunc("/tag", handlers.HandleTagAdd).Methods("PUT")
	router.HandleFunc("/tag/{id:[0-9]+}", handlers.HandleTagEdit).Methods("POST")
	router.HandleFunc("/tag/{id:[0-9]+}", handlers.HandleTagDelete).Methods("DELETE")
}
