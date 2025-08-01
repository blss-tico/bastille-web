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
