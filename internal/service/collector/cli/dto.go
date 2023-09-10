package cli

type DTO struct {
	PluginType    string      `json:"pluginType"`
	PluginVersion string      `json:"pluginVersion"`
	CliType       string      `json:"cliType"`
	CliVersion    string      `json:"cliVersion"`
	OSName        string      `json:"osName"`
	IdeType       string      `json:"ideType,omitempty"`
	IdeVersion    string      `json:"ideVersion,omitempty"`
	Events        []DTOEvents `json:"events"`
}

type DTOEvents struct {
	Uid            string         `json:"uid"`
	CreatedAt      string         `json:"createdAt"`
	Type           string         `json:"type"`
	Project        string         `json:"project,omitempty"`
	ProjectBaseDir string         `json:"projectBaseDir,omitempty"`
	Language       string         `json:"language,omitempty"`
	Target         string         `json:"target,omitempty"`
	Branch         string         `json:"branch,omitempty"`
	Timezone       string         `json:"timezone"`
	Params         map[string]any `json:"params,omitempty"`
}
