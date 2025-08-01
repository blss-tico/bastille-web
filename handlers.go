package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func bootstrapHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("bootstrapHandler")

	var data bootstrapModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleBootstrap(data.Options, data.ReleaseTemplate, data.UpdateArch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}
