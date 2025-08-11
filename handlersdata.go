package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type HandlersData struct {
	bl Bastille
}

// bootstrap
// @Summary bootstrap command
// @Description The bootstrap sub-command is used to download and extract releases and templates for use with Bastille containers. A valid release is needed before containers can be created. Templates are optional but are managed in the same manner.
// @Tags bootstrap
// @Accept  json
// @Produce  text/plain
// @Param  bootstrap  body	bootstrapModel  true  "bootstrap"
// @Success 200 {object} string
// @Router /bootstrap [post]
func (hd *HandlersData) bootstrap(w http.ResponseWriter, r *http.Request) {
	log.Println("bootstrapHandler")

	var data bootstrapModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.bootstrap(data.Options, data.ReleaseTemplate, data.UpdateArch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// clone
// @Summary clone command
// @Description Clone/duplicate an existing jail to a new jail.
// @Tags clone
// @Accept  json
// @Produce  text/plain
// @Param  clone  body	cloneModel  true  "clone"
// @Success 200 {object} string
// @Router /clone [post]
func (hd *HandlersData) clone(w http.ResponseWriter, r *http.Request) {
	log.Println("cloneHandler")

	var data cloneModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.clone(data.Options, data.Target, data.Newname, data.Ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// cmd
// @Summary cmd command
// @Description Execute commands inside targeted jail(s).
// @Tags cmd
// @Accept  json
// @Produce  text/plain
// @Param  cmd  body	cmdModel  true  "cmd"
// @Success 200 {object} string
// @Router /cmd [post]
func (hd *HandlersData) cmd(w http.ResponseWriter, r *http.Request) {
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

	result, err := hd.bl.cmd(data.Options, data.Target, commands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// config
// @Summary config command
// @Description Get, set or remove properties from targeted jail(s).
// @Tags config
// @Accept  json
// @Produce  text/plain
// @Param  config  body	configModel  true  "config"
// @Success 200 {object} string
// @Router /config [post]
func (hd *HandlersData) config(w http.ResponseWriter, r *http.Request) {
	log.Println("configHandler")

	var data configModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.config(data.Options, data.Target, data.Action, data.Property, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// console
// @Summary console command
// @Description Launch a login shell into the jail. Default is password- less root login.
// @Tags console
// @Accept  json
// @Produce  text/plain
// @Param  console  body	consoleModel  true  "config"
// @Success 200 {object} string
// @Router /console [post]
func (hd *HandlersData) console(w http.ResponseWriter, r *http.Request) {
	log.Println("consoleHandler")

	var data consoleModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.console(data.Options, data.Target, data.User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// convert
// @Summary convert command
// @Description Convert a thin jail to a thick jail. Convert a thick jail to a custom release.
// @Tags convert
// @Accept  json
// @Produce  text/plain
// @Param  convert  body	convertModel  true  "config"
// @Success 200 {object} string
// @Router /convert [post]
func (hd *HandlersData) convert(w http.ResponseWriter, r *http.Request) {
	log.Println("convertHandler")

	var data convertModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.convert(data.Options, data.Target, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// cp
// @Summary cp command
// @Description Copy files from host to jail(s).
// @Tags cp
// @Accept  json
// @Produce  text/plain
// @Param  cp  body	cpModel  true  "cp"
// @Success 200 {object} string
// @Router /cp [post]
func (hd *HandlersData) cp(w http.ResponseWriter, r *http.Request) {
	log.Println("cpHandler")

	var data cpModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.cp(data.Options, data.Target, data.Hostpath, data.Jailpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// create
// @Summary create command
// @Description Create a jail uning any available bootstrapped release. To create a jail, simply provide a name, bootstrapped release, and IP address.
// @Tags create
// @Accept  json
// @Produce  text/plain
// @Param  create  body	createModel  true  "create"
// @Success 200 {object} string
// @Router /create [post]
func (hd *HandlersData) create(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandler")

	var data createModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.create(data.Options, data.Name, data.Release, data.Ip, data.Iface, data.Ipip, data.Value, data.Vlanid, data.Zfsoptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// destroy
// @Summary destroy command
// @Description Destroy jails or releases.
// @Tags destroy
// @Accept  json
// @Produce  text/plain
// @Param  destroy  body	destroyModel  true  "destroy"
// @Success 200 {object} string
// @Router /destroy [post]
func (hd *HandlersData) destroy(w http.ResponseWriter, r *http.Request) {
	log.Println("destroyHandler")

	var data destroyModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.destroy(data.Options, data.JailRelease)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// edit
// @Summary edit command
// @Description Edit jail config files.
// @Tags edit
// @Accept  json
// @Produce  text/plain
// @Param  edit  body	editModel  true  "edit"
// @Success 200 {object} string
// @Router /edit [post]
func (hd *HandlersData) edit(w http.ResponseWriter, r *http.Request) {
	log.Println("editHandler")

	var data editModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.edit(data.Options, data.Target, data.File)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// etcupdate
// @Summary etcupdate command
// @Description This command will update the contents of /etc inside a jail. It should be run after a jail upgrade.
// @Tags etcupdate
// @Accept  json
// @Produce  text/plain
// @Param  etcupdate  body	etcupdateModel  true  "etcupdate"
// @Success 200 {object} string
// @Router /etcupdate [post]
func (hd *HandlersData) etcupdate(w http.ResponseWriter, r *http.Request) {
	log.Println("etcupdateHandler")

	var data etcupdateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.etcupdate(data.Options, data.Bootstraptarget, data.Action, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// export
// @Summary export command
// @Description Exporting a container creates an archive or image that can be sent to a different machine to be imported later. These exported archives can be used as container backups.
// @Tags export
// @Accept  json
// @Produce  text/plain
// @Param  export  body	exportModel  true  "export"
// @Success 200 {object} string
// @Router /export [post]
func (hd *HandlersData) export(w http.ResponseWriter, r *http.Request) {
	log.Println("exportHandler")

	var data exportModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.export(data.Options, data.Target, data.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// htop
// @Summary htop command
// @Description This command runs htop in the targeted jail. Requires htop to be installed in the jail.
// @Tags htop
// @Accept  json
// @Produce  text/plain
// @Param  htop  body	htopModel  true  "htop"
// @Success 200 {object} string
// @Router /htop [post]
func (hd *HandlersData) htop(w http.ResponseWriter, r *http.Request) {
	log.Println("htopHandler")

	var data htopModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.htop(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// import
// @Summary import command
// @Description Import a jail backup image or archive.
// @Tags import
// @Accept  json
// @Produce  text/plain
// @Param  import  body	importModel  true  "import"
// @Success 200 {object} string
// @Router /import [post]
func (hd *HandlersData) imporT(w http.ResponseWriter, r *http.Request) {
	log.Println("importHandler")

	var data importModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.imporT(data.Options, data.File, data.Release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// jcp
// @Summary jcp command
// @Description Copy files from jail to jail(s).
// @Tags jcp
// @Accept  json
// @Produce  text/plain
// @Param  jcp  body	jcpModel  true  "jcp"
// @Success 200 {object} string
// @Router /jcp [post]
func (hd *HandlersData) jcp(w http.ResponseWriter, r *http.Request) {
	log.Println("jcpHandler")

	var data jcpModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.jcp(data.Options, data.Sourcejail, data.Jailpath, data.Destjail, data.Jailpath2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// limits
// @Summary limits command
// @Description Set resourse limits for targeted jail(s).
// @Tags limits
// @Accept  json
// @Produce  text/plain
// @Param  limits  body	limitsModel  true  "limits"
// @Success 200 {object} string
// @Router /limits [post]
func (hd *HandlersData) limits(w http.ResponseWriter, r *http.Request) {
	log.Println("limitsHandler")

	var data limitsModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.limits(data.Options, data.Target, data.Action, data.Option, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// list
// @Summary list command
// @Description List jails, ports, releases, templates, logs, limits, exports and imports and much more managed by bastille.
// @Tags list
// @Accept  json
// @Produce  text/plain
// @Param  list  body	listModel  true  "list"
// @Success 200 {object} string
// @Router /list [post]
func (hd *HandlersData) list(w http.ResponseWriter, r *http.Request) {
	log.Println("listHandler")

	var data listModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.list(data.Options, data.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// migrate
// @Summary migrate command
// @Description The migrate sub-command allows migrating the targeted jail(s) to another remote system.
// @Tags migrate
// @Accept  json
// @Produce  text/plain
// @Param  migrate  body	migrateModel  true  "migrate"
// @Success 200 {object} string
// @Router /migrate [post]
func (hd *HandlersData) migrate(w http.ResponseWriter, r *http.Request) {
	log.Println("migrateHandler")

	var data migrateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.migrate(data.Options, data.Target, data.Remote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// mount
// @Summary mount command
// @Description To mount storage within the container use bastille mount.
// @Tags mount
// @Accept  json
// @Produce  text/plain
// @Param  mount  body	mountModel  true  "mount"
// @Success 200 {object} string
// @Router /mount [post]
func (hd *HandlersData) mount(w http.ResponseWriter, r *http.Request) {
	log.Println("mountHandler")

	var data mountModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.mount(data.Options, data.Target, data.Hostpath, data.Jailpath, data.Filesystemtype, data.Option, data.Dump, data.Passnumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// network
// @Summary network command
// @Description Add or remove interfaces to existing jails.
// @Tags network
// @Accept  json
// @Produce  text/plain
// @Param  network  body	networkModel  true  "network"
// @Success 200 {object} string
// @Router /network [post]
func (hd *HandlersData) network(w http.ResponseWriter, r *http.Request) {
	log.Println("networkHandler")

	var data networkModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.network(data.Options, data.Target, data.Action, data.Iface, data.Ip, data.Vlanid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// pkg
// @Summary pkg command
// @Description Manage binary packages inside jails.
// @Tags pkg
// @Accept  json
// @Produce  text/plain
// @Param  pkg  body	pkgModel  true  "pkg"
// @Success 200 {object} string
// @Router /pkg [post]
func (hd *HandlersData) pkg(w http.ResponseWriter, r *http.Request) {
	log.Println("pkgHandler")

	var data pkgModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	arg := strings.Split(data.Arg, " ")
	if len(arg) == 0 {
		http.Error(w, "args is not found", http.StatusBadRequest)
		return
	}

	result, err := hd.bl.pkg(data.Options, data.Target, arg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// rcp
// @Summary rcp command
// @Description This command allows copying files from jail to host.
// @Tags rcp
// @Accept  json
// @Produce  text/plain
// @Param  rcp  body	rcpModel  true  "rcp"
// @Success 200 {object} string
// @Router /rcp [post]
func (hd *HandlersData) rcp(w http.ResponseWriter, r *http.Request) {
	log.Println("rcpHandler")

	var data rcpModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.rcp(data.Options, data.Target, data.Jailpath, data.Hostpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// rdr
// @Summary rdr command
// @Description Bastille rdr allows you to configure dynamic rdr rules for your containers without modifying pf.conf.
// @Tags rdr
// @Accept  json
// @Produce  text/plain
// @Param  rdr  body	rdrModel  true  "rdr"
// @Success 200 {object} string
// @Router /rdr [post]
func (hd *HandlersData) rdr(w http.ResponseWriter, r *http.Request) {
	log.Println("rdrHandler")

	var data rdrModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.rdr(data.Options, data.Optionsarg, data.Target, data.Action, data.Hostport, data.Jailport, data.Log, data.Logopts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

// rename
// @Summary rename command
// @Description Rename a jail.
// @Tags rename
// @Accept  json
// @Produce  text/plain
// @Param  rename  body	renameModel  true  "rename"
// @Success 200 {object} string
// @Router /rename [post]
func (hd *HandlersData) rename(w http.ResponseWriter, r *http.Request) {
	log.Println("renameHandler")

	var data renameModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.rename(data.Options, data.Target, data.Newname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) restart(w http.ResponseWriter, r *http.Request) {
	log.Println("restartHandler")

	var data restartModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.restart(data.Options, data.Target, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) service(w http.ResponseWriter, r *http.Request) {
	log.Println("serviceHandler")

	var data serviceModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.service(data.Options, data.Target, data.Servicename, data.Args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) setup(w http.ResponseWriter, r *http.Request) {
	log.Println("setupHandler")

	var data setupModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.setup(data.Options, data.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) start(w http.ResponseWriter, r *http.Request) {
	log.Println("startHandler")

	var data startModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.start(data.Options, data.Target, data.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) stop(w http.ResponseWriter, r *http.Request) {
	log.Println("stopHandler")

	var data stopModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.stop(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) sysrc(w http.ResponseWriter, r *http.Request) {
	log.Println("sysrcHandler")

	var data sysrcModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.sysrc(data.Options, data.Target, data.Args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) tags(w http.ResponseWriter, r *http.Request) {
	log.Println("tagsHandler")

	var data tagsModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.tags(data.Options, data.Target, data.Action, data.Tgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) template(w http.ResponseWriter, r *http.Request) {
	log.Println("templateHandler")

	var data templateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Println(data)
	result, err := hd.bl.template(data.Options, data.Target, data.Action, data.Template)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) top(w http.ResponseWriter, r *http.Request) {
	log.Println("topHandler")

	var data topModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.top(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) umount(w http.ResponseWriter, r *http.Request) {
	log.Println("umountHandler")

	var data umountModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.umount(data.Options, data.Target, data.Jailpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) update(w http.ResponseWriter, r *http.Request) {
	log.Println("updateHandler")

	var data updateModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.update(data.Options, data.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) upgrade(w http.ResponseWriter, r *http.Request) {
	log.Println("upgradeHandler")

	var data upgradeModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.upgrade(data.Options, data.Target, data.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) verify(w http.ResponseWriter, r *http.Request) {
	log.Println("verifyHandler")

	var data verifyModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.verify(data.Options, data.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}

func (hd *HandlersData) zfs(w http.ResponseWriter, r *http.Request) {
	log.Println("zfsHandler")

	var data zfsModel
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := hd.bl.zfs(data.Options, data.Target, data.Action, data.Tag, data.Key, data.Value, data.Pooldataset, data.Jailpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOkWithJSONUtil(w, result)
}
