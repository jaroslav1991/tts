package sender

import (
	"time"

	"github.com/jaroslav1991/tts/internal/model"
)

// todo make test
func NewRemoteRequestDTOFromDataModels(models []model.DataModel) RemoteRequestDTO {
	result := make(RemoteRequestDTO, len(models))
	for i, item := range models {
		result[i] = RemoteRequestDTOItem{
			Program:          item.PluginInfo.Program,
			Duration:         item.PluginInfo.Duration,
			PathProject:      item.PluginInfo.PathProject,
			CurrentGitBranch: item.AggregatorInfo.CurrentGitBranch,
		}
	}
	return result
}

type RemoteRequestDTO []RemoteRequestDTOItem

type RemoteRequestDTOItem struct {
	Program          string        `json:"program"`
	Duration         time.Duration `json:"duration"`
	PathProject      string        `json:"pathProject"`
	CurrentGitBranch string        `json:"currentGitBranch"`
}
