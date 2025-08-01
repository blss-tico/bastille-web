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

func convertHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("convertHandler")

	var data convertModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleConvert(data.Options, data.Target, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func cpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("cpHandler")

	var data cpModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleCp(data.Options, data.Target, data.Hostpath, data.Jailpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandler")

	var data createModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleCreate(data.Options, data.Name, data.Release, data.Ip, data.Iface, data.Ipip, data.Value, data.Vlanid, data.Zfsoptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func destroyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("destroyHandler")

	var data destroyModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleDestroy(data.Options, data.JailRelease)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func etcupdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("etcupdateHandler")

	var data etcupdateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleEtcupdate(data.Options, data.Bootstraptarget, data.Action, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func exportHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("exportHandler")

	var data exportModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleExport(data.Options, data.Target, data.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func htopHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("htopHandler")

	var data htopModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleHtop(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func importHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("importHandler")

	var data importModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleImport(data.Options, data.File, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func jcpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("jcpHandler")

	var data jcpModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := bastilleJcp(data.Options, data.Sourcejail, data.Jailpath, data.Destjail, data.Jailpath2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}
