package handlers

import (
	"github.com/gorilla/mux"
	"io"
	"katinka/database"
	"katinka/services/helpers"
	"log"
	"net/http"
	"strconv"
)

func HandleImageHit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		internalError(err, w)
		return
	}

	err = database.RegisterHit(id)
	if err != nil {
		internalError(err, w)
		return
	}

	response(jsonResponse{
		Action:  "image_hit",
		Success: true,
	}, w)
}

func HandleUploadPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(500 * 1024 * 1024 * 1024)
	if err != nil {
		log.Printf("Handler upload: - Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	files := r.MultipartForm.File["upload_images"]
	for _, fh := range files {
		file, err := fh.Open()
		defer file.Close()

		if err != nil {
			internalError(err, w)
			return
		}

		cont, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Handler upload: - Error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		err = database.SaveImage(cont, fh.Filename)
		if err != nil {
			log.Printf("Handler upload: - Error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	response(jsonResponse{
		Action:  "files_upload",
		Success: true,
		Data:    "/upload",
	}, w)
}

func HandleGetImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Handler upload: - Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	image, err := database.GetImage(id)

	if err != nil {
		log.Printf("Handler upload: - Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if image.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(image.ImageBlob)
}

func HandleImageSearch(w http.ResponseWriter, r *http.Request) {
	tagsIds, err := helpers.StringToNumArray(r.FormValue("tags_ids"))
	if err != nil {
		internalError(err, w)
		return
	}

	pageParam := r.FormValue("page")
	if pageParam == "" {
		pageParam = "0"
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		internalError(err, w)
		return
	}

	images, err := database.SearchImages(tagsIds, page)

	response(jsonResponse{
		Action:  "image_search",
		Success: true,
		Data:    images,
	}, w)
}
