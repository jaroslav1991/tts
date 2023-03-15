package sender

import "github.com/jaroslav1991/tts/internal/model"

func NewRemoteRequestDTOFromDataModels(models []model.DataModel) RemoteRequestDTO {
	result := make(RemoteRequestDTO, len(models))
	for i, item := range models {

		var events []DTOEvents
		for _, event := range item.PluginInfo.Events {
			dtoEvent := DTOEvents{
				CreatedAt:      event.CreatedAt,
				Type:           event.Type,
				Project:        event.Project,
				ProjectBaseDir: event.ProjectBaseDir,
				Language:       event.Language,
				Target:         event.Target,
				Branch:         event.Branch,
				Params:         event.Params,
			}

			if eventBranch, ok := item.AggregatorInfo.GitBranchesByProjectBaseDir[event.ProjectBaseDir]; ok {
				dtoEvent.Branch = eventBranch
			}

			events = append(events, dtoEvent)
		}

		result[i] = RemoteRequestDTOItem{
			Uid:           item.PluginInfo.Uid,
			PluginType:    item.PluginInfo.PluginType,
			PluginVersion: item.PluginInfo.PluginVersion,
			IdeType:       item.PluginInfo.IdeType,
			IdeVersion:    item.PluginInfo.IdeVersion,
			Events:        events,
		}
	}
	return result
}

type RemoteRequestDTO []RemoteRequestDTOItem

type RemoteRequestDTOItem struct {
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
