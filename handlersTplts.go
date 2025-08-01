package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func homeHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("homeHandlerTplt")

	type SysInfo struct {
		Hostname        string `json:"hostname"`
		Arch            string `json:"arch"`
		Platform        string `json:"platform"`
		Osrelease       string `json:"osrelease"`
		Totalmemory     string `json:"totalmemory"`
		Availablememory string `json:"availablememory"`
		BastilleVersion string `json:"bastilleversion"`
	}

	osinf, _ := osInfo()
	posinf := strings.Split(osinf, " ")

	mminf, _ := osMemInfo()
	re := regexp.MustCompile(`\d+`)
	pmminf := re.FindAllString(mminf, -1)

	bstv, _ := bastilleVersion()

	sysinfo := SysInfo{
		Hostname:        posinf[1],
		Arch:            posinf[7],
		Platform:        posinf[0],
		Osrelease:       posinf[2],
		Totalmemory:     pmminf[0],
		Availablememory: pmminf[1],
		BastilleVersion: bstv,
	}

	result, err := bastilleListAllJson()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	type List struct {
		Jid       string `json:"jid"`
		Boot      string `json:"boot"`
		Prio      string `json:"prio"`
		State     string `json:"state"`
		Ip        string `json:"ip address"`
		Published string `json:"published ports"`
		Hostname  string `json:"hostname"`
		Release   string `json:"release"`
		Path      string `json:"path"`
	}

	var list []List
	err = json.Unmarshal([]byte(result), &list)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	type HomeData struct {
		Title     string
		Data      bastilleModel
		SysData   SysInfo
		JailsData []List
	}
	data := HomeData{Title: "Home", Data: bastille, SysData: sysinfo, JailsData: list}
	renderTemplateUtil(w, "home.html", data)
}

func helpHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("helpHandlerTplt")
	data := templateModel{Title: "Help", Data: bastille}
	renderTemplateUtil(w, "help.html", data)
}

func contactHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("contactHandlerTplt")
	data := templateModel{Title: "Contact", Data: bastille}
	renderTemplateUtil(w, "contact.html", data)
}

func bootstrapHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("bootstrapHandlerTplt")
	data := templateModel{Title: "Bootstrap", Data: bastille}
	renderTemplateUtil(w, "bootstrap.html", data)
}
