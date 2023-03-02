package sender

import "github.com/jaroslav1991/tts/internal/model"

func NewRemoteRequestDTOFromDataModels(models []model.DataModel) RemoteRequestDTO {
	result := make(RemoteRequestDTO, len(models))
	for i, item := range models {

		var events []DTOEvents
		for _, event := range item.PluginInfo.Events {
			dtoEvent := DTOEvents{
				Uid:       event.Uid,
				CreatedAt: event.CreatedAt,
				Type:      event.Type,
				Project:   event.Project,
				Language:  event.Language,
				Target:    event.Target,
				Params:    event.Params,
			}

			if eventBranch, ok := item.AggregatorInfo.GitBranchesByEventUID[event.Uid]; ok {
				dtoEvent.Branch = eventBranch
			}

			events = append(events, dtoEvent)
		}

		result[i] = RemoteRequestDTOItem{
			PluginType:    item.PluginInfo.PluginType,
			PluginVersion: item.PluginInfo.PluginVersion,
			CliType:       item.PluginInfo.CliType,
			CliVersion:    item.PluginInfo.CliVersion,
			DeviceName:    item.PluginInfo.DeviceName,
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
