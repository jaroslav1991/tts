package sender

// todo make test
//func NewRemoteRequestDTOFromDataModels(models []model.DataModel) RemoteRequestDTO {
//	result := make(RemoteRequestDTO, len(models))
//	for i, item := range models {
//		result[i] = RemoteRequestDTOItem{
//			PluginType:    item.PluginInfo.PluginType,
//			PluginVersion: item.PluginInfo.PluginVersion,
//			CliType:       item.PluginInfo.CliType,
//			CliVersion:    item.PluginInfo.CliVersion,
//			DeviceName:    item.PluginInfo.DeviceName,
//			Events: DTOEvents{
//				Uid:       item.PluginInfo.Events.Uid,
//				CreatedAt: item.PluginInfo.Events.CreatedAt,
//				Type:      item.PluginInfo.Events.Type,
//				Project:   item.PluginInfo.Events.Project,
//				Language:  item.PluginInfo.Events.Language,
//				Target:    item.PluginInfo.Events.Target,
//				Branch:    &item.AggregatorInfo.CurrentGitBranch,
//				Params:    *item.PluginInfo.Events.Params,
//			},
//		}
//	}
//	return result
//}

type RemoteRequestDTO []RemoteRequestDTOItem

type RemoteRequestDTOItem struct {
	PluginType    string    `json:"pluginType"`
	PluginVersion string    `json:"pluginVersion"`
	CliType       string    `json:"cliType"`
	CliVersion    string    `json:"cliVersion"`
	DeviceName    string    `json:"deviceName,omitempty"`
	Events        DTOEvents `json:"events"`
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
