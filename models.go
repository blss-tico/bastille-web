package main

const addrModel = "127.0.0.1:80"

type optionsModel struct {
	Sflag string
	Lflag string
	Text  string
}

type bastilleCommandsModel struct {
	Command     string         `json:"command"`
	Description string         `json:"description"`
	Options     []optionsModel `json:"options"`
	Fields      []string       `json:"fields"`
	Help        string         `json:"help"`
	HelpUrl     string         `json:"helpUrl"`
}

type bastilleModel struct {
	Software string                  `json:"software"`
	Options  []string                `json:"options"`
	Help     string                  `json:"help"`
	Commands []bastilleCommandsModel `json:"commands"`
}

type templatesModel struct {
	Title string
	Data  bastilleModel
}

type bootstrapModel struct {
	Options         string `json:"options,omitempty" example:"-x" format:"string"`
	ReleaseTemplate string `json:"release|template" example:"14.3-RELEASE" format:"string"`
	UpdateArch      string `json:"update|arch,omitempty" example:"--i386" format:"string"`
}

type cloneModel struct {
	Options string `json:"options,omitempty" example:"-a|-l|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string"`
	Newname string `json:"new_name" example:"new_jail_name" format:"string"`
	Ip      string `json:"ip" example:"n.n.n.n" format:"string"`
}

type cmdModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string"`
	Command string `json:"command" example:"ls -l" format:"string"`
}

type configModel struct {
	Options  string `json:"options,omitempty" example:"-x" format:"string"`
	Target   string `json:"target" example:"jail_target" format:"string"`
	Action   string `json:"action" example:"get|set" format:"string"`
	Property string `json:"property" example:"ip4.addr" format:"string"`
	Value    string `json:"value,omitempty" example:"depends on property" format:"string"`
}

type consoleModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string"`
	User    string `json:"user,omitempty" example:"root" format:"string"`
}

type convertModel struct {
	Options string `json:"options,omitempty" example:"-a|-y|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string"`
	Release string `json:"release,omitempty" example:"myrelease" format:"string"`
}

type cpModel struct {
	Options  string `json:"options,omitempty" example:"-q|-x" format:"string"`
	Target   string `json:"target" example:"jail_target" format:"string"`
	Hostpath string `json:"hostpath" example:"/host/path" format:"string"`
	Jailpath string `json:"jailpath" example:"/jail/path" format:"string"`
}

type createModel struct {
	Options    string `json:"options,omitempty" example:"-B|-C|-D|-E|-g|-L|-M|-n|--no-validate|--no-boot|-p|-T|-V|-v|-x|-Z" format:"string"`
	Name       string `json:"name" example:"jail_name" format:"string"`
	Release    string `json:"release" example:"14.3-RELEASE" format:"string"`
	Ip         string `json:"ip" example:"n.n.n.n" format:"string"`
	Iface      string `json:"iface,omitempty" example:"bastille0" format:"string"`
	Ipip       string `json:"ipip,omitempty" example:"n.n.n.n,i.i.i.i" format:"string"`
	Value      string `json:"value,omitempty" example:"" format:"string"`
	Vlanid     string `json:"vlanid,omitempty" example:"vlan10" format:"string"`
	Zfsoptions string `json:"zfsoptions,omitempty" example:"" format:"string"`
}

type destroyModel struct {
	Options     string `json:"options,omitempty" example:"-a|-c|-f|-y|-x" format:"string"`
	JailRelease string `json:"jail|release" example:"jail_target or release" format:"string"`
}

type editModel struct {
	Options string `json:"options,omitempty" example:"-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string"`
	File    string `json:"file,omitempty" example:"file_name" format:"string"`
}

type etcupdateModel struct {
	Options         string `json:"options,omitempty" example:"-d|-f|-x" format:"string"`
	Bootstraptarget string `json:"bootstrap|target" example:"jail_name|bootstrap" format:"string"`
	Action          string `json:"action,omitempty" example:"diff|resolve|update" format:"string"`
	Release         string `json:"release,omitempty" example:"14.3-RELEASE" format:"string"`
}

type exportModel struct {
	Options string `json:"options,omitempty" example:"-a|--gz|-r|-s|--tgz|--txz|-v|--xz|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string"`
	Path    string `json:"path" example:"/host/path" format:"string"`
}

type htopModel struct {
	Options string `json:"options,omitempty" example:"-a|-x" format:"string"`
	Target  string `json:"target" example:"jail_target" format:"string"`
}

type importModel struct {
	Options string `json:"options,omitempty" example:"-f|-M|-v|-x" format:"string"`
	File    string `json:"file" example:"/path/to/archive.file" format:"string"`
	Release string `json:"release,omitempty" example:"release_name" format:"string"`
}

type jcpModel struct {
	Options    string `json:"options,omitempty" example:"-q|-x" format:"string"`
	Sourcejail string `json:"source_jail" example:"source_jail" format:"string"`
	Jailpath   string `json:"jail_path" example:"/source_jail/path" format:"string"`
	Destjail   string `json:"dest_jail" example:"dest_jail" format:"string"`
	Jailpath2  string `json:"jail_path2" example:"/dest_jail/path" format:"string"`
}

type limitsModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Action  string `json:"action"`
	Option  string `json:"option,omitempty"`
	Value   string `json:"value,omitempty"`
}

type listModel struct {
	Options string `json:"options,omitempty"`
	Action  string `json:"action"`
}

type migrateModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Remote  string `json:"remote"`
}

type mountModel struct {
	Options        string `json:"options,omitempty"`
	Target         string `json:"target"`
	Hostpath       string `json:"host_path"`
	Jailpath       string `json:"jail_path"`
	Filesystemtype string `json:"filesystem_type"`
	Option         string `json:"option"`
	Dump           string `json:"dump"`
	Passnumber     string `json:"pass_number"`
}

type networkModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Action  string `json:"action"`
	Iface   string `json:"interface,omitempty"`
	Ip      string `json:"ip,omitempty"`
}

type pkgModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Arg     string `json:"args"`
}

type rcpModel struct {
	Options  string `json:"options,omitempty"`
	Target   string `json:"target"`
	Jailpath string `json:"jail_path"`
	Hostpath string `json:"host_path"`
}

type rdrModel struct {
	Options    string `json:"options,omitempty"`
	Optionsarg string `json:"options_arg,omitempty"`
	Target     string `json:"target"`
	Action     string `json:"action"`
	Hostport   string `json:"host_port"`
	Jailport   string `json:"jail_port"`
	Log        string `json:"log,omitempty"`
	Logopts    string `json:"logopts,omitempty"`
}

type renameModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Newname string `json:"new_name"`
}

type restartModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Value   string `json:"value,omitempty"`
}

type serviceModel struct {
	Options     string `json:"options,omitempty"`
	Target      string `json:"target"`
	Servicename string `json:"service_name"`
	Args        string `json:"args"`
}

type setupModel struct {
	Options string `json:"options,omitempty"`
	Action  string `json:"action"`
}

type startModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Value   string `json:"value,omitempty"`
}

type stopModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
}

type sysrcModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Args    string `json:"args"`
}

type tagsModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Action  string `json:"action"`
	Tgs     string `json:"tags"`
}

type templateModel struct {
	Options  string `json:"options,omitempty"`
	Target   string `json:"target"`
	Action   string `json:"action,omitempty"`
	Template string `json:"template"`
}

type topModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
}

type umountModel struct {
	Options  string `json:"options,omitempty"`
	Target   string `json:"target"`
	Jailpath string `json:"jail_path"`
}

type updateModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
}

type upgradeModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Action  string `json:"NEW_RELEASE|install"`
}

type verifyModel struct {
	Options string `json:"options,omitempty"`
	Action  string `json:"RELEASE|TEMPLATE"`
}

type zfsModel struct {
	Options     string `json:"options,omitempty"`
	Target      string `json:"target"`
	Action      string `json:"action"`
	Tag         string `json:"tag,omitempty"`
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	Pooldataset string `json:"pool/dataset,omitempty"`
	Jailpath    string `json:"/jail/path,omitempty"`
}
