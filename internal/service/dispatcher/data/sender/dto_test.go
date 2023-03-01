package sender

import (
	"encoding/json"
	"testing"
)

func TestOmitEmpty(t *testing.T) {
	marshal, _ := json.Marshal(RemoteRequestDTOItem{})
	t.Log(string(marshal))
}

//func TestNewRemoteRequestDTOFromDataModels(t *testing.T) {
//	type testCase struct {
//		models   []model.DataModel
//		expected RemoteRequestDTO
//	}
//	var info model.PluginInfo
//
//	testCases := map[string]testCase{
//		"empty models": {
//			models:   nil,
//			expected: RemoteRequestDTO{},
//		},
//		"one model": {
//			models: []model.DataModel{
//				{
//					PluginInfo: model.PluginInfo{
//						PluginType:    "1",
//						PluginVersion: "1",
//						CliType:       "1",
//						CliVersion:    "1",
//						DeviceName:    "1",
//						Events: model.Events{
//							Uid:       "1",
//							CreatedAt: "1",
//							Type:      "1",
//							Project:   "1",
//							Language:  "1",
//							Target:    "1",
//							Branch:    "master",
//							Params:    nil,
//						},
//					},
//					AggregatorInfo: model.AggregatorInfo{
//						CurrentGitBranch: "master",
//					},
//				},
//			},
//			expected: RemoteRequestDTO{
//				{
//					PluginType:    "1",
//					PluginVersion: "1",
//					CliType:       "1",
//					CliVersion:    "1",
//					DeviceName:    "1",
//					Events: DTOEvents{
//						Uid:       "1",
//						CreatedAt: "1",
//						Type:      "1",
//						Project:   "1",
//						Language:  "1",
//						Target:    "1",
//						Branch:    &info.Events.Branch,
//						Params:    nil,
//					},
//				},
//			},
//		},
//		"some model": {
//			models: []model.DataModel{
//				{
//					PluginInfo: model.PluginInfo{
//						PluginType:    "1",
//						PluginVersion: "1",
//						CliType:       "1",
//						CliVersion:    "1",
//						DeviceName:    "1",
//						Events: model.Events{
//							Uid:       "1",
//							CreatedAt: "1",
//							Type:      "1",
//							Project:   "1",
//							Language:  "1",
//							Target:    "1",
//							Branch:    "master",
//							Params:    nil,
//						},
//					},
//					AggregatorInfo: model.AggregatorInfo{
//						CurrentGitBranch: "master",
//					},
//				},
//				{
//					PluginInfo: model.PluginInfo{
//						PluginType:    "2",
//						PluginVersion: "2",
//						CliType:       "2",
//						CliVersion:    "2",
//						DeviceName:    "2",
//						Events: model.Events{
//							Uid:       "2",
//							CreatedAt: "2",
//							Type:      "2",
//							Project:   "2",
//							Language:  "2",
//							Target:    "2",
//							Branch:    "master2",
//							Params:    nil,
//						},
//					},
//					AggregatorInfo: model.AggregatorInfo{
//						CurrentGitBranch: "master2",
//					},
//				},
//			},
//			expected: RemoteRequestDTO{
//				{
//					PluginType:    "1",
//					PluginVersion: "1",
//					CliType:       "1",
//					CliVersion:    "1",
//					DeviceName:    "1",
//					Events: DTOEvents{
//						Uid:       "1",
//						CreatedAt: "1",
//						Type:      "1",
//						Project:   "1",
//						Language:  "1",
//						Target:    "1",
//						Branch:    &info.Events.Branch,
//						Params:    nil,
//					},
//				},
//				{
//					PluginType:    "2",
//					PluginVersion: "2",
//					CliType:       "2",
//					CliVersion:    "2",
//					DeviceName:    "2",
//					Events: DTOEvents{
//						Uid:       "2",
//						CreatedAt: "2",
//						Type:      "2",
//						Project:   "2",
//						Language:  "2",
//						Target:    "2",
//						Branch:    &info.Events.Branch,
//						Params:    nil,
//					},
//				},
//			},
//		},
//	}
//
//	for name, tc := range testCases {
//		t.Run(name, func(t *testing.T) {
//			assert.Equal(t, tc.expected, NewRemoteRequestDTOFromDataModels(tc.models))
//		})
//	}
//}
