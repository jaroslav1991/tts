package sender

import (
	"github.com/jaroslav1991/tts/internal/model"
)

func NewRemoteRequestDTOFromDataModels(models []model.DataModel) RemoteRequestDTOItem {
	var result RemoteRequestDTOItem
	for _, item := range models {

		var events []DTOEvents
		for j, event := range item.PluginInfo.Events {
			dtoEvent := DTOEvents{
				Id:             event.Id,
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

			if dtoEvent.Id == "" {
				dtoEvent.Id = item.AggregatorInfo.Id[j]
			}

			events = append(events, dtoEvent)
		}

		result.Events = append(result.Events, events...)

	}
	return result
}

type RemoteRequestDTOItem struct {
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
