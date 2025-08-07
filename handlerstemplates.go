package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type HandlersTemplates struct {
	bl Bastille
}

func (ht *HandlersTemplates) home(w http.ResponseWriter, r *http.Request) {
	log.Println("homeHandlersTemplates")

	type SysInfo struct {
		Hostname        string `json:"hostname"`
		Arch            string `json:"arch"`
		Platform        string `json:"platform"`
		Osrelease       string `json:"osrelease"`
		Totalmemory     string `json:"totalmemory"`
		Availablememory string `json:"availablememory"`
		BastilleVersion string `json:"bastilleversion"`
	}

	osinf, _ := infoOsUtil()
	posinf := strings.Split(osinf, " ")

	mminf, _ := memInfoOsUtil()
	re := regexp.MustCompile(`\d+`)
	pmminf := re.FindAllString(mminf, -1)

	bstv, _ := ht.bl.bastilleVersion()

	sysinfo := SysInfo{
		Hostname:        posinf[1],
		Arch:            posinf[7],
		Platform:        posinf[0],
		Osrelease:       posinf[2],
		Totalmemory:     pmminf[0],
		Availablememory: pmminf[1],
		BastilleVersion: bstv,
	}

	result, err := ht.bl.list("-j", "all")
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
		Addr      string
	}

	data := HomeData{
		Title:     "home",
		Data:      bastille,
		SysData:   sysinfo,
		JailsData: list,
		Addr:      addrModel,
	}

	renderTemplateUtil(w, "home.html", data)
}

func (ht *HandlersTemplates) help(w http.ResponseWriter, r *http.Request) {
	log.Println("helpHandlersTemplates")
	data := templatesModel{Title: "help", Data: bastille}
	renderTemplateUtil(w, "help.html", data)
}

func (ht *HandlersTemplates) contact(w http.ResponseWriter, r *http.Request) {
	log.Println("contactHandlersTemplates")

	type ContactModel struct {
		Title      string
		Data       bastilleModel
		Email      string
		Githubpers string
		Githubproj string
	}

	const email = "blss-tico@gmail.com"
	const githubpers = "https://github.com/blss-tico"
	const githubproj = "https://github.com/blss-tico/bastille-web"
	data := ContactModel{
		Title:      "contact",
		Data:       bastille,
		Email:      email,
		Githubpers: githubpers,
		Githubproj: githubproj,
	}

	renderTemplateUtil(w, "contact.html", data)
}

func (ht *HandlersTemplates) bootstrap(w http.ResponseWriter, r *http.Request) {
	log.Println("bootstrapHandlersTemplates")
	data := templatesModel{Title: "bootstrap", Data: bastille}
	renderTemplateUtil(w, "bootstrap.html", data)
}

func (ht *HandlersTemplates) clone(w http.ResponseWriter, r *http.Request) {
	log.Println("cloneHandlersTemplates")
	data := templatesModel{Title: "clone", Data: bastille}
	renderTemplateUtil(w, "clone.html", data)
}

func (ht *HandlersTemplates) cmd(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlersTemplates")
	data := templatesModel{Title: "cmd", Data: bastille}
	renderTemplateUtil(w, "cmd.html", data)
}

func (ht *HandlersTemplates) config(w http.ResponseWriter, r *http.Request) {
	log.Println("cmdHandlersTemplates")
	data := templatesModel{Title: "config", Data: bastille}
	renderTemplateUtil(w, "config.html", data)
}

func (ht *HandlersTemplates) console(w http.ResponseWriter, r *http.Request) {
	log.Println("consoleHandlersTemplates")
	data := templatesModel{Title: "console", Data: bastille}
	renderTemplateUtil(w, "console.html", data)
}

func (ht *HandlersTemplates) convert(w http.ResponseWriter, r *http.Request) {
	log.Println("convertHandlersTemplates")
	data := templatesModel{Title: "convert", Data: bastille}
	renderTemplateUtil(w, "convert.html", data)
}

func (ht *HandlersTemplates) cp(w http.ResponseWriter, r *http.Request) {
	log.Println("cpHandlersTemplates")
	data := templatesModel{Title: "cp", Data: bastille}
	renderTemplateUtil(w, "cp.html", data)
}

func (ht *HandlersTemplates) create(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandlersTemplates")
	data := templatesModel{Title: "create", Data: bastille}
	renderTemplateUtil(w, "create.html", data)
}

func (ht *HandlersTemplates) destroy(w http.ResponseWriter, r *http.Request) {
	log.Println("destroyHandlersTemplates")
	data := templatesModel{Title: "destroy", Data: bastille}
	renderTemplateUtil(w, "destroy.html", data)
}

func (ht *HandlersTemplates) edit(w http.ResponseWriter, r *http.Request) {
	log.Println("editHandlersTemplates")
	data := templatesModel{Title: "edit", Data: bastille}
	renderTemplateUtil(w, "edit.html", data)
}

func (ht *HandlersTemplates) etcupdate(w http.ResponseWriter, r *http.Request) {
	log.Println("etcupdateHandlersTemplates")
	data := templatesModel{Title: "etcupdate", Data: bastille}
	renderTemplateUtil(w, "etcupdate.html", data)
}

func (ht *HandlersTemplates) export(w http.ResponseWriter, r *http.Request) {
	log.Println("exportHandlersTemplates")
	data := templatesModel{Title: "export", Data: bastille}
	renderTemplateUtil(w, "export.html", data)
}

func (ht *HandlersTemplates) htop(w http.ResponseWriter, r *http.Request) {
	log.Println("htopHandlersTemplates")
	data := templatesModel{Title: "htop", Data: bastille}
	renderTemplateUtil(w, "htop.html", data)
}

func (ht *HandlersTemplates) imporT(w http.ResponseWriter, r *http.Request) {
	log.Println("importHandlersTemplates")
	data := templatesModel{Title: "import", Data: bastille}
	renderTemplateUtil(w, "import.html", data)
}

func (ht *HandlersTemplates) jcp(w http.ResponseWriter, r *http.Request) {
	log.Println("jcpHandlersTemplates")
	data := templatesModel{Title: "jcp", Data: bastille}
	renderTemplateUtil(w, "jcp.html", data)
}

func (ht *HandlersTemplates) limits(w http.ResponseWriter, r *http.Request) {
	log.Println("limitsHandlersTemplates")
	data := templatesModel{Title: "limits", Data: bastille}
	renderTemplateUtil(w, "limits.html", data)
}

func (ht *HandlersTemplates) list(w http.ResponseWriter, r *http.Request) {
	log.Println("listHandlersTemplates")
	data := templatesModel{Title: "list", Data: bastille}
	renderTemplateUtil(w, "list.html", data)
}

func (ht *HandlersTemplates) migrate(w http.ResponseWriter, r *http.Request) {
	log.Println("migrateHandlersTemplates")
	data := templatesModel{Title: "migrate", Data: bastille}
	renderTemplateUtil(w, "migrate.html", data)
}

func (ht *HandlersTemplates) mount(w http.ResponseWriter, r *http.Request) {
	log.Println("mountHandlersTemplates")
	data := templatesModel{Title: "mount", Data: bastille}
	renderTemplateUtil(w, "mount.html", data)
}

func (ht *HandlersTemplates) network(w http.ResponseWriter, r *http.Request) {
	log.Println("networkHandlersTemplates")
	data := templatesModel{Title: "network", Data: bastille}
	renderTemplateUtil(w, "network.html", data)
}

func (ht *HandlersTemplates) pkg(w http.ResponseWriter, r *http.Request) {
	log.Println("pkgHandlersTemplates")
	data := templatesModel{Title: "pkg", Data: bastille}
	renderTemplateUtil(w, "pkg.html", data)
}

func (ht *HandlersTemplates) rcp(w http.ResponseWriter, r *http.Request) {
	log.Println("rcpHandlersTemplates")
	data := templatesModel{Title: "rcp", Data: bastille}
	renderTemplateUtil(w, "rcp.html", data)
}

func (ht *HandlersTemplates) rdr(w http.ResponseWriter, r *http.Request) {
	log.Println("rdrHandlersTemplates")
	data := templatesModel{Title: "rdr", Data: bastille}
	renderTemplateUtil(w, "rdr.html", data)
}

func (ht *HandlersTemplates) rename(w http.ResponseWriter, r *http.Request) {
	log.Println("renameHandlersTemplates")
	data := templatesModel{Title: "rename", Data: bastille}
	renderTemplateUtil(w, "rename.html", data)
}

func (ht *HandlersTemplates) restart(w http.ResponseWriter, r *http.Request) {
	log.Println("restartHandlersTemplates")
	data := templatesModel{Title: "restart", Data: bastille}
	renderTemplateUtil(w, "restart.html", data)
}

func (ht *HandlersTemplates) service(w http.ResponseWriter, r *http.Request) {
	log.Println("serviceHandlersTemplates")
	data := templatesModel{Title: "service", Data: bastille}
	renderTemplateUtil(w, "service.html", data)
}

func (ht *HandlersTemplates) setup(w http.ResponseWriter, r *http.Request) {
	log.Println("setupHandlersTemplates")
	data := templatesModel{Title: "setup", Data: bastille}
	renderTemplateUtil(w, "setup.html", data)
}

func (ht *HandlersTemplates) start(w http.ResponseWriter, r *http.Request) {
	log.Println("startHandlersTemplates")
	data := templatesModel{Title: "start", Data: bastille}
	renderTemplateUtil(w, "start.html", data)
}

func (ht *HandlersTemplates) stop(w http.ResponseWriter, r *http.Request) {
	log.Println("stopHandlersTemplates")
	data := templatesModel{Title: "stop", Data: bastille}
	renderTemplateUtil(w, "stop.html", data)
}

func (ht *HandlersTemplates) sysrc(w http.ResponseWriter, r *http.Request) {
	log.Println("sysrcHandlersTemplates")
	data := templatesModel{Title: "sysrc", Data: bastille}
	renderTemplateUtil(w, "sysrc.html", data)
}

func (ht *HandlersTemplates) tags(w http.ResponseWriter, r *http.Request) {
	log.Println("tagsHandlersTemplates")
	data := templatesModel{Title: "tags", Data: bastille}
	renderTemplateUtil(w, "tags.html", data)
}

func (ht *HandlersTemplates) template(w http.ResponseWriter, r *http.Request) {
	log.Println("templateHandlersTemplates")
	data := templatesModel{Title: "template", Data: bastille}
	renderTemplateUtil(w, "template.html", data)
}

func (ht *HandlersTemplates) top(w http.ResponseWriter, r *http.Request) {
	log.Println("topHandlersTemplates")
	data := templatesModel{Title: "top", Data: bastille}
	renderTemplateUtil(w, "top.html", data)
}

func (ht *HandlersTemplates) umount(w http.ResponseWriter, r *http.Request) {
	log.Println("umountHandlersTemplates")
	data := templatesModel{Title: "umount", Data: bastille}
	renderTemplateUtil(w, "umount.html", data)
}

func (ht *HandlersTemplates) update(w http.ResponseWriter, r *http.Request) {
	log.Println("updateHandlersTemplates")
	data := templatesModel{Title: "update", Data: bastille}
	renderTemplateUtil(w, "update.html", data)
}

func (ht *HandlersTemplates) upgrade(w http.ResponseWriter, r *http.Request) {
	log.Println("upgradeHandlersTemplates")
	data := templatesModel{Title: "upgrade", Data: bastille}
	renderTemplateUtil(w, "upgrade.html", data)
}

func (ht *HandlersTemplates) verify(w http.ResponseWriter, r *http.Request) {
	log.Println("verifyHandlersTemplates")
	data := templatesModel{Title: "verify", Data: bastille}
	renderTemplateUtil(w, "verify.html", data)
}
