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
