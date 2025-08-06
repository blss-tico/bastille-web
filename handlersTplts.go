package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var templates *template.Template

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
		Ip        string
	}
	data := HomeData{Title: "home", Data: bastille, SysData: sysinfo, JailsData: list, Ip: IpAddrModel}
	renderTemplateUtil(w, "home.html", data)
}

func helpHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("helpHandlerTplt")
	data := templatesModel{Title: "help", Data: bastille}
	renderTemplateUtil(w, "help.html", data)
}

func contactHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("contactHandlerTplt")

	type ContactModel struct {
		Title      string
		Data       bastilleModel
		Email      string
		Githubpers string
		Githubproj string
	}

	email := "blss-tico@gmail.com"
	githubpers := "https://github.com/blss-tico"
	githubproj := "https://github.com/blss-tico/bastille-web"
	data := ContactModel{Title: "contact", Data: bastille, Email: email, Githubpers: githubpers, Githubproj: githubproj}
	renderTemplateUtil(w, "contact.html", data)
}

func bootstrapHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("bootstrapHandlerTplt")
	data := templatesModel{Title: "bootstrap", Data: bastille}
	renderTemplateUtil(w, "bootstrap.html", data)
}

func cloneHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("cloneHandlerTplt")
	data := templatesModel{Title: "clone", Data: bastille}
	renderTemplateUtil(w, "clone.html", data)
}

func cmdHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlerTplt")
	data := templatesModel{Title: "cmd", Data: bastille}
	renderTemplateUtil(w, "cmd.html", data)
}

func configHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlerTplt")
	data := templatesModel{Title: "config", Data: bastille}
	renderTemplateUtil(w, "config.html", data)
}

func convertHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("convertHandlerTplt")
	data := templatesModel{Title: "convert", Data: bastille}
	renderTemplateUtil(w, "convert.html", data)
}

func cpHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("cpHandlerTplt")
	data := templatesModel{Title: "cp", Data: bastille}
	renderTemplateUtil(w, "cp.html", data)
}

func createHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandlerTplt")
	data := templatesModel{Title: "create", Data: bastille}
	renderTemplateUtil(w, "create.html", data)
}

func destroyHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("destroyHandlerTplt")
	data := templatesModel{Title: "destroy", Data: bastille}
	renderTemplateUtil(w, "destroy.html", data)
}

func etcupdateHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("etcupdateHandlerTplt")
	data := templatesModel{Title: "etcupdate", Data: bastille}
	renderTemplateUtil(w, "etcupdate.html", data)
}

func exportHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("exportHandlerTplt")
	data := templatesModel{Title: "export", Data: bastille}
	renderTemplateUtil(w, "export.html", data)
}

func htopHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("htopHandlerTplt")
	data := templatesModel{Title: "htop", Data: bastille}
	renderTemplateUtil(w, "htop.html", data)
}

func importHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("importHandlerTplt")
	data := templatesModel{Title: "import", Data: bastille}
	renderTemplateUtil(w, "import.html", data)
}

func jcpHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("jcpHandlerTplt")
	data := templatesModel{Title: "jcp", Data: bastille}
	renderTemplateUtil(w, "jcp.html", data)
}

func limitsHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("limitsHandlerTplt")
	data := templatesModel{Title: "limits", Data: bastille}
	renderTemplateUtil(w, "limits.html", data)
}

func listHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("listHandlerTplt")
	data := templatesModel{Title: "list", Data: bastille}
	renderTemplateUtil(w, "list.html", data)
}

func migrateHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("migrateHandlerTplt")
	data := templatesModel{Title: "migrate", Data: bastille}
	renderTemplateUtil(w, "migrate.html", data)
}

func mountHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("mountHandlerTplt")
	data := templatesModel{Title: "mount", Data: bastille}
	renderTemplateUtil(w, "mount.html", data)
}

func networkHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("networkHandlerTplt")
	data := templatesModel{Title: "network", Data: bastille}
	renderTemplateUtil(w, "network.html", data)
}

func pkgHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("pkgHandlerTplt")
	data := templatesModel{Title: "pkg", Data: bastille}
	renderTemplateUtil(w, "pkg.html", data)
}

func rcpHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("rcpHandlerTplt")
	data := templatesModel{Title: "rcp", Data: bastille}
	renderTemplateUtil(w, "rcp.html", data)
}

func renameHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("renameHandlerTplt")
	data := templatesModel{Title: "rename", Data: bastille}
	renderTemplateUtil(w, "rename.html", data)
}

func restartHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("restartHandlerTplt")
	data := templatesModel{Title: "restart", Data: bastille}
	renderTemplateUtil(w, "restart.html", data)
}

func serviceHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("serviceHandlerTplt")
	data := templatesModel{Title: "service", Data: bastille}
	renderTemplateUtil(w, "service.html", data)
}

func setupHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("setupHandlerTplt")
	data := templatesModel{Title: "setup", Data: bastille}
	renderTemplateUtil(w, "setup.html", data)
}

func startHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("startHandlerTplt")
	data := templatesModel{Title: "start", Data: bastille}
	renderTemplateUtil(w, "start.html", data)
}

func stopHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("stopHandlerTplt")
	data := templatesModel{Title: "stop", Data: bastille}
	renderTemplateUtil(w, "stop.html", data)
}

func sysrcHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("sysrcHandlerTplt")
	data := templatesModel{Title: "sysrc", Data: bastille}
	renderTemplateUtil(w, "sysrc.html", data)
}

func tagsHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("tagsHandlerTplt")
	data := templatesModel{Title: "tags", Data: bastille}
	renderTemplateUtil(w, "tags.html", data)
}

func templateHandlerTplt(w http.ResponseWriter, r *http.Request) {
	log.Println("templateHandlerTplt")
	data := templatesModel{Title: "template", Data: bastille}
	renderTemplateUtil(w, "template.html", data)
}
