package cli

type DTO struct {
	Events []DTOEvents `json:"events"`
}

type DTOEvents struct {
	Id             string         `json:"id"`
	CreatedAt      string         `json:"createdAt"`
	Type           string         `json:"type"`
	Project        string         `json:"project,omitempty"`
	ProjectBaseDir string         `json:"projectBaseDir,omitempty"`
	Language       string         `json:"language,omitempty"`
	Target         string         `json:"target,omitempty"`
	Branch         string         `json:"branch,omitempty"`
	Timezone       string         `json:"timezone,omitempty"`
	Params         map[string]any `json:"params,omitempty"`
}
