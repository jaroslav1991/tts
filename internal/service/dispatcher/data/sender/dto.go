package sender

import "github.com/jaroslav1991/tts/internal/model"

func NewRemoteRequestDTOFromDataModels(models []model.DataModel) RemoteRequestDTO {
	result := make(RemoteRequestDTO, len(models))
	for i, item := range models {

		var events []DTOEvents
		for _, event := range item.PluginInfo.Events {
			dtoEvent := DTOEvents{
				Uid:            item.AggregatorInfo.Uid,
				CreatedAt:      event.CreatedAt,
				Type:           event.Type,
				Project:        event.Project,
				ProjectBaseDir: event.ProjectBaseDir,
				Language:       event.Language,
				Target:         event.Target,
				Branch:         event.Branch,
				Timezone:       event.Timezone,
				Params:         event.Params,
			}

			if eventBranch, ok := item.AggregatorInfo.GitBranchesByProjectBaseDir[event.ProjectBaseDir]; ok {
				dtoEvent.Branch = eventBranch
			}

			events = append(events, dtoEvent)
		}

		result[i] = RemoteRequestDTOItem{
			PluginType:    item.PluginInfo.PluginType,
			PluginVersion: item.PluginInfo.PluginVersion,
			CliType:       item.PluginInfo.CliType,
			CliVersion:    item.PluginInfo.CliVersion,
			OSName:        item.AggregatorInfo.OSName,
			IdeType:       item.PluginInfo.IdeType,
			IdeVersion:    item.PluginInfo.IdeVersion,
			Events:        events,
		}
	}
	return result
}

type RemoteRequestDTO []RemoteRequestDTOItem

type RemoteRequestDTOItem struct {
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
