package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type jsonResponse struct {
	Action  string      `json:"action"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func view(tplName string, tplData interface{}, w http.ResponseWriter) {

	//TODO Templates can be cached in prod version
	tpl, err := template.New(tplName).ParseFiles("./templates/" + tplName)
	if err != nil {
		log.Printf("Template %v: - Can not get template: %v", tplName, err)
	}

	tpl, err = tpl.ParseGlob("./templates/elements/*.gohtml")
	if err != nil {
		log.Printf("Err %v: - Can not get templates elements", err)
	}

	err = tpl.Execute(w, tplData)
	if err != nil {
		log.Printf("Template %v - Template exec failed: %v", tplName, err)
	}
}

func internalError(err error, w http.ResponseWriter) {
	message := fmt.Sprintf("Error: %v", err)
	log.Print(message)
	w.WriteHeader(http.StatusInternalServerError)

	return
}

func response(data jsonResponse, w http.ResponseWriter) {
	strBytes, _ := json.Marshal(data)
	w.Write(strBytes)
}
