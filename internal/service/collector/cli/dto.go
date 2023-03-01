package cli

type DTO struct {
	PluginType    string      `json:"pluginType"`
	PluginVersion string      `json:"pluginVersion"`
	CliType       string      `json:"cliType"`
	CliVersion    string      `json:"cliVersion"`
	DeviceName    string      `json:"deviceName,omitempty"`
	Events        []DTOEvents `json:"events"`
}

type DTOEvents struct {
	Uid       string         `json:"uid"`
	CreatedAt string         `json:"createdAt"`
	Type      string         `json:"type"`
	Project   string         `json:"project,omitempty"`
	Language  string         `json:"language,omitempty"`
	Target    string         `json:"target,omitempty"`
	Branch    string         `json:"branch,omitempty"`
	Params    map[string]any `json:"params,omitempty"`
}
