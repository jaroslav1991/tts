package sender

import (
	"testing"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNewRemoteRequestDTOFromDataModels(t *testing.T) {
	type testCase struct {
		models   []model.DataModel
		expected RemoteRequestDTO
	}

	testCases := map[string]testCase{
		"empty models": {
			models:   nil,
			expected: RemoteRequestDTO{},
		},
		"one model": {
			models: []model.DataModel{
				{
					PluginInfo: model.PluginInfo{
						PluginType:    "1",
						PluginVersion: "1",
						CliType:       "windowsOS",
						CliVersion:    "1.1.0",
						OSName:        "",
						IdeType:       "1",
						IdeVersion:    "1",
						Events: []model.Events{
							{
								Id:             "",
								CreatedAt:      "1",
								Type:           "1",
								Project:        "1",
								ProjectBaseDir: "some-base",
								Language:       "1",
								Target:         "1",
								Branch:         "",
								Timezone:       "",
								Params:         nil,
							},
						},
					},
					AggregatorInfo: model.AggregatorInfo{
						GitBranchesByProjectBaseDir: map[string]string{
							"some-base": "master",
						},
						OSName: "windows",
						Id:     "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
					},
				},
			},
			expected: RemoteRequestDTO{
				{
					PluginType:    "1",
					PluginVersion: "1",
					CliType:       "windowsOS",
					CliVersion:    "1.1.0",
					OSName:        "windows",
					IdeType:       "1",
					IdeVersion:    "1",
					Events: []DTOEvents{
						{
							Id:             "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
							CreatedAt:      "1",
							Type:           "1",
							Project:        "1",
							ProjectBaseDir: "some-base",
							Language:       "1",
							Target:         "1",
							Branch:         "master",
							Timezone:       "",
							Params:         nil,
						},
					},
				},
			},
		},
		"some model": {
			models: []model.DataModel{
				{
					PluginInfo: model.PluginInfo{
						PluginType:    "1",
						PluginVersion: "1",
						CliType:       "windowsOS1",
						CliVersion:    "1.1.01",
						OSName:        "",
						IdeType:       "1",
						IdeVersion:    "1",
						Events: []model.Events{
							{
								Id:             "",
								CreatedAt:      "1",
								Type:           "1",
								Project:        "1",
								ProjectBaseDir: "some-base",
								Language:       "1",
								Target:         "1",
								Branch:         "",
								Timezone:       "1",
								Params:         nil,
							},
						},
					},
					AggregatorInfo: model.AggregatorInfo{
						GitBranchesByProjectBaseDir: map[string]string{
							"some-base1": "master1",
						},
						OSName: "windows1",
						Id:     "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
					},
				},
				{
					PluginInfo: model.PluginInfo{
						PluginType:    "2",
						PluginVersion: "2",
						CliType:       "windowsOS2",
						CliVersion:    "1.1.02",
						OSName:        "",
						IdeType:       "2",
						IdeVersion:    "2",
						Events: []model.Events{
							{
								Id:             "",
								CreatedAt:      "2",
								Type:           "2",
								Project:        "2",
								ProjectBaseDir: "some-base2",
								Language:       "2",
								Target:         "2",
								Branch:         "",
								Timezone:       "2",
								Params:         nil,
							},
						},
					},
					AggregatorInfo: model.AggregatorInfo{
						GitBranchesByProjectBaseDir: map[string]string{
							"some-base2": "master2",
						},
						OSName: "windows2",
						Id:     "a6ac8ef0-28e2-4b6e-8568-aa8934f53c85",
					},
				},
			},
			expected: RemoteRequestDTO{
				{
					PluginType:    "1",
					PluginVersion: "1",
					CliType:       "windowsOS1",
					CliVersion:    "1.1.01",
					OSName:        "windows1",
					IdeType:       "1",
					IdeVersion:    "1",
					Events: []DTOEvents{
						{
							Id:             "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
							CreatedAt:      "1",
							Type:           "1",
							Project:        "1",
							ProjectBaseDir: "some-base",
							Language:       "1",
							Target:         "1",
							Timezone:       "1",
							Params:         nil,
						},
					},
				},
				{
					PluginType:    "2",
					PluginVersion: "2",
					CliType:       "windowsOS2",
					CliVersion:    "1.1.02",
					OSName:        "windows2",
					IdeType:       "2",
					IdeVersion:    "2",
					Events: []DTOEvents{
						{
							Id:             "a6ac8ef0-28e2-4b6e-8568-aa8934f53c85",
							CreatedAt:      "2",
							Type:           "2",
							Project:        "2",
							ProjectBaseDir: "some-base2",
							Language:       "2",
							Target:         "2",
							Branch:         "master2",
							Timezone:       "2",
							Params:         nil,
						},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewRemoteRequestDTOFromDataModels(tc.models))
		})
	}
}
