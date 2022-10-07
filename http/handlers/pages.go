package handlers

import (
	"github.com/gorilla/mux"
	"katinka/database"
	"net/http"
	"strconv"
)

func HandleAuthPage(w http.ResponseWriter, r *http.Request) {
	view("auth.gohtml", nil, w)
}

func HandleMainPage(w http.ResponseWriter, r *http.Request) {

	totalHits, err := database.GetTotalHits()

	if err != nil {
		internalError(err, w)
		return
	}

	lastUploaded, err := database.GetLastUploaded(10)

	if err != nil {
		internalError(err, w)
		return
	}

	view("index.gohtml", struct {
		AppName      string
		TotalHits    int
		LastUploaded []database.Image
	}{
		AppName:      "KaTinKa",
		TotalHits:    totalHits,
		LastUploaded: lastUploaded,
	}, w)
}

func HandleUploadPage(w http.ResponseWriter, r *http.Request) {
	view("upload.gohtml", struct{}{}, w)
}

func HandleSearchPage(w http.ResponseWriter, r *http.Request) {
	tags, err := database.GetTagsList()

	if err != nil {
		internalError(err, w)
		return
	}

	view("search.gohtml", struct {
		Tags []database.Tag
	}{
		Tags: tags,
	}, w)
}

func HandleViewRandPage(w http.ResponseWriter, r *http.Request) {
	images, err := database.GetRandImages(10)

	if err != nil {
		internalError(err, w)
		return
	}

	tplData := struct {
		Title  string
		Images []database.Image
	}{
		Title:  "KaTinKa Random",
		Images: images,
	}

	view("view_rand.gohtml", tplData, w)
}

func HandleViewImagePage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		internalError(err, w)
		return
	}

	image, err := database.GetImage(id)
	if err != nil {
		internalError(err, w)
		return
	}

	if image.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tags, err := database.GetTagsByImageId(id)

	if err != nil {
		internalError(err, w)
		return
	}

	tplData := struct {
		Title string
		Image database.Image
		Tags  []database.Tag
		Next  int
	}{
		Title: "Image edit",
		Image: image,
		Tags:  tags,
		Next:  image.Id + 1,
	}

	view("image_page.gohtml", tplData, w)
}
