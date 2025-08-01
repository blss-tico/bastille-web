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

func cloneHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("cloneHandlerTplt")
	data := templateModel{Title: "Contact", Data: bastille}
	renderTemplateUtil(w, "clone.html", data)
}

func cmdHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlerTplt")
	data := templateModel{Title: "Cmd", Data: bastille}
	renderTemplateUtil(w, "cmd.html", data)
}

func configHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlerTplt")
	data := templateModel{Title: "Cmd", Data: bastille}
	renderTemplateUtil(w, "config.html", data)
}

func convertHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("convertHandlerTplt")
	data := templateModel{Title: "Convert", Data: bastille}
	renderTemplateUtil(w, "convert.html", data)
}

func cpHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("cpHandlerTplt")
	data := templateModel{Title: "Cp", Data: bastille}
	renderTemplateUtil(w, "cp.html", data)
}

func createHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandlerTplt")
	data := templateModel{Title: "Create", Data: bastille}
	renderTemplateUtil(w, "create.html", data)
}

func destroyHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("destroyHandlerTplt")
	data := templateModel{Title: "Destroy", Data: bastille}
	renderTemplateUtil(w, "destroy.html", data)
}

func etcupdateHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("etcupdateHandlerTplt")
	data := templateModel{Title: "Etcupdate", Data: bastille}
	renderTemplateUtil(w, "etcupdate.html", data)
}

func exportHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("exportHandlerTplt")
	data := templateModel{Title: "Export", Data: bastille}
	renderTemplateUtil(w, "export.html", data)
}

func htopHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("htopHandlerTplt")
	data := templateModel{Title: "Htop", Data: bastille}
	renderTemplateUtil(w, "htop.html", data)
}

func importHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("importHandlerTplt")
	data := templateModel{Title: "Import", Data: bastille}
	renderTemplateUtil(w, "import.html", data)
}

func jcpHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("jcpHandlerTplt")
	data := templateModel{Title: "Jcp", Data: bastille}
	renderTemplateUtil(w, "jcp.html", data)
}

func limitsHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("limitsHandlerTplt")
	data := templateModel{Title: "Limits", Data: bastille}
	renderTemplateUtil(w, "limits.html", data)
}

func listHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("listHandlerTplt")
	data := templateModel{Title: "List", Data: bastille}
	renderTemplateUtil(w, "list.html", data)
}

func migrateHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("migrateHandlerTplt")
	data := templateModel{Title: "Migrate", Data: bastille}
	renderTemplateUtil(w, "migrate.html", data)
}
