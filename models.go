package main

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
}

type bastilleModel struct {
	Software string                  `json:"software"`
	Options  []string                `json:"options"`
	Help     string                  `json:"help"`
	Commands []bastilleCommandsModel `json:"commands"`
}

type templateModel struct {
	Title string
	Data  bastilleModel
}

type bootstrapModel struct {
	Options         string `json:"options,omitempty"`
	ReleaseTemplate string `json:"release|template"`
	UpdateArch      string `json:"update|arch,omitempty"`
}

type cloneModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Newname string `json:"new_name"`
	Ip      string `json:"ip"`
}

type cmdModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Command string `json:"command"`
}

type configModel struct {
	Options  string `json:"options,omitempty"`
	Target   string `json:"target"`
	Action   string `json:"action"`
	Property string `json:"property"`
	Value    string `json:"value,omitempty"`
}

type convertModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Release string `json:"release,omitempty"`
}

type cpModel struct {
	Options  string `json:"options,omitempty"`
	Target   string `json:"target"`
	Hostpath string `json:"hostpath"`
	Jailpath string `json:"jailpath"`
}

type createModel struct {
	Options    string `json:"options,omitempty"`
	Name       string `json:"name"`
	Release    string `json:"release"`
	Ip         string `json:"ip"`
	Iface      string `json:"iface,omitempty"`
	Ipip       string `json:"ipip,omitempty"`
	Value      string `json:"value,omitempty"`
	Vlanid     string `json:"vlanid,omitempty"`
	Zfsoptions string `json:"zfsoptions,omitempty"`
}

type destroyModel struct {
	Options     string `json:"options,omitempty"`
	JailRelease string `json:"jail|release"`
}

type etcupdateModel struct {
	Options         string `json:"options,omitempty"`
	Bootstraptarget string `json:"bootstrap|target"`
	Action          string `json:"action,omitempty"`
	Release         string `json:"release,omitempty"`
}

type exportModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Path    string `json:"path"`
}

type htopModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
}

type importModel struct {
	Options string `json:"options,omitempty"`
	File    string `json:"file"`
	Release string `json:"release,omitempty"`
}

type jcpModel struct {
	Options    string `json:"options,omitempty"`
	Sourcejail string `json:"source_jail"`
	Jailpath   string `json:"jail_path"`
	Destjail   string `json:"dest_jail"`
	Jailpath2  string `json:"jail_path2"`
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

type startModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
	Value   string `json:"value,omitempty"`
}

type stopModel struct {
	Options string `json:"options,omitempty"`
	Target  string `json:"target"`
}
