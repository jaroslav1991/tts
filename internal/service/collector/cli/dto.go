package cli

type DTO struct {
	Uid           string      `json:"uid"`
	PluginType    string      `json:"PluginType"`
	PluginVersion string      `json:"pluginVersion"`
	IdeType       string      `json:"ideType,omitempty"`
	IdeVersion    string      `json:"ideVersion,omitempty"`
	Events        []DTOEvents `json:"events"`
}

type DTOEvents struct {
	CreatedAt      string         `json:"createdAt"`
	Type           string         `json:"type"`
	Project        string         `json:"project,omitempty"`
	ProjectBaseDir string         `json:"projectBaseDir,omitempty"`
	Language       string         `json:"language,omitempty"`
	Target         string         `json:"target,omitempty"`
	Branch         string         `json:"branch,omitempty"`
	Params         map[string]any `json:"params,omitempty"`
}
