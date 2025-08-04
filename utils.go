package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func renderTemplateUtil(w http.ResponseWriter, name string, data any) {
	log.Println("renderTemplate")

	tmplt := "./templates/" + name
	files := []string{
		"./templates/base.html",
		"./templates/partials/navbar.html",
		"./templates/partials/sidebar.html",
	}
	files = append(files, tmplt)

	var err error
	templates, err = template.ParseFiles(files...)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	err = templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func respondOkWithJSONUtil(w http.ResponseWriter, payload string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if payload != "" {
		response := map[string]string{"msg": payload}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]string{"msg": "command executed"}
	json.NewEncoder(w).Encode(response)
}

func RespondErrorWithJSONUtil(w http.ResponseWriter, code int, payload string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	response := map[string]string{"msg": "with errors", "err": payload}
	json.NewEncoder(w).Encode(response)
}
