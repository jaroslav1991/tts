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
						Uid:           "qwerty123",
						PluginType:    "1",
						PluginVersion: "1",
						IdeType:       "1",
						IdeVersion:    "1",
						Events: []model.Events{
							{
								CreatedAt:      "1",
								Type:           "1",
								Project:        "1",
								ProjectBaseDir: "some-base",
								Language:       "1",
								Target:         "1",
								Branch:         "",
								Params:         nil,
							},
						},
					},
					AggregatorInfo: model.AggregatorInfo{
						GitBranchesByProjectBaseDir: map[string]string{
							"some-base": "master",
						},
					},
				},
			},
			expected: RemoteRequestDTO{
				{
					Uid:           "qwerty123",
					PluginType:    "1",
					PluginVersion: "1",
					IdeType:       "1",
					IdeVersion:    "1",
					Events: []DTOEvents{
						{
							CreatedAt:      "1",
							Type:           "1",
							Project:        "1",
							ProjectBaseDir: "some-base",
							Language:       "1",
							Target:         "1",
							Branch:         "master",
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
						Uid:           "qwerty123",
						PluginType:    "1",
						PluginVersion: "1",
						IdeType:       "1",
						IdeVersion:    "1",
						Events: []model.Events{
							{
								CreatedAt:      "1",
								Type:           "1",
								Project:        "1",
								ProjectBaseDir: "some-base",
								Language:       "1",
								Target:         "1",
								Branch:         "",
								Params:         nil,
							},
						},
					},
					AggregatorInfo: model.AggregatorInfo{},
				},
				{
					PluginInfo: model.PluginInfo{
						Uid:           "qwerty234",
						PluginType:    "2",
						PluginVersion: "2",
						IdeType:       "2",
						IdeVersion:    "2",
						Events: []model.Events{
							{
								CreatedAt:      "2",
								Type:           "2",
								Project:        "2",
								ProjectBaseDir: "some-base2",
								Language:       "2",
								Target:         "2",
								Branch:         "",
								Params:         nil,
							},
						},
					},
					AggregatorInfo: model.AggregatorInfo{
						GitBranchesByProjectBaseDir: map[string]string{
							"some-base2": "master2",
						},
					},
				},
			},
			expected: RemoteRequestDTO{
				{
					Uid:           "qwerty123",
					PluginType:    "1",
					PluginVersion: "1",
					IdeType:       "1",
					IdeVersion:    "1",
					Events: []DTOEvents{
						{
							CreatedAt:      "1",
							Type:           "1",
							Project:        "1",
							ProjectBaseDir: "some-base",
							Language:       "1",
							Target:         "1",
							Params:         nil,
						},
					},
				},
				{
					Uid:           "qwerty234",
					PluginType:    "2",
					PluginVersion: "2",
					IdeType:       "2",
					IdeVersion:    "2",
					Events: []DTOEvents{
						{
							CreatedAt:      "2",
							Type:           "2",
							Project:        "2",
							ProjectBaseDir: "some-base2",
							Language:       "2",
							Target:         "2",
							Branch:         "master2",
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
