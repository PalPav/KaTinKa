package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"katinka/database"
	"net/http"
	"strconv"
)

func HandleTagSearch(w http.ResponseWriter, r *http.Request) {

	tags, err := database.SearchTags(r.FormValue("query"))

	if err != nil {
		internalError(err, w)
		return
	}

	response(jsonResponse{
		Action:  "tag_search",
		Success: true,
		Data:    tags,
	}, w)
}

type tagImageData struct {
	ImageId int `json:"image_id"`
	TagId   int `json:"tag_id"`
}

func HandleTagAssign(w http.ResponseWriter, r *http.Request) {

	var requestData tagImageData
	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		internalError(err, w)
		return
	}

	tag, err := database.GetTag(requestData.TagId)
	if err != nil {
		internalError(err, w)
		return
	}

	if tag.Id <= 0 { //TODO add proper http code )
		internalError(fmt.Errorf("invalid tag id"), w)
		return
	}

	err = database.AssignTag(requestData.ImageId, requestData.TagId)
	if err != nil {
		internalError(err, w)
		return
	}

	response(jsonResponse{
		Action:  "tag_assign",
		Success: true,
		Data:    tag,
	}, w)
}

func HandleTagDetach(w http.ResponseWriter, r *http.Request) {
	var requestData tagImageData
	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		internalError(err, w)
		return
	}

	err = database.DetachTag(requestData.ImageId, requestData.TagId)
	if err != nil {
		internalError(err, w)
		return
	}

	response(jsonResponse{
		Action:  "tag_detach",
		Success: true,
	}, w)
}

func HandleTagAdd(w http.ResponseWriter, r *http.Request) {

	var requestData database.Tag
	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		internalError(err, w)
		return
	}

	tag, err := database.AddTag(requestData.Name)
	if err != nil {
		internalError(err, w)
		return
	}

	response(jsonResponse{
		Action:  "tag_detach",
		Success: true,
		Data:    tag,
	}, w)
}

func HandleTagEdit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		internalError(err, w)
		return
	}
	if id <= 0 {
		internalError(fmt.Errorf("invalid tag id"), w)
		return
	}

	var tag database.Tag
	err = json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		internalError(err, w)
		return
	}

	tag.Id = id
	err = database.EditTag(tag)
	if err != nil {
		internalError(err, w)
		return
	}

	response(jsonResponse{
		Action:  "tag_edit",
		Success: true,
		Data:    tag,
	}, w)
}

func HandleTagDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		internalError(err, w)
		return
	}
	if id <= 0 {
		internalError(fmt.Errorf("invalid tag id"), w)
		return
	}

	err = database.RemoveTag(id)
	if err != nil {
		internalError(err, w)
		return
	}

	response(jsonResponse{
		Action:  "tag_remove",
		Success: true,
	}, w)
}
