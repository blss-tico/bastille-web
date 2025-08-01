package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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

func cloneHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("cloneHandler")

	var data cloneModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleClone(data.Options, data.Target, data.Newname, data.Ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandler")

	var data cmdModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	commands := strings.Split(data.Command, " ")
	if len(commands) == 0 {
		http.Error(w, "command not found", http.StatusBadRequest)
		return
	}

	result, err := bastilleCmd(data.Options, data.Target, commands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("configHandler")

	var data configModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleConfig(data.Options, data.Target, data.Action, data.Property, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}
