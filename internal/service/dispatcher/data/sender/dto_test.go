package sender

import (
	"testing"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNewRemoteRequestDTOFromDataModels(t *testing.T) {
	type testCase struct {
		models   []model.DataModel
		expected RemoteRequestDTOItem
	}

	testCases := map[string]testCase{
		"empty models": {
			models:   nil,
			expected: RemoteRequestDTOItem{},
		},
		"one model": {
			models: []model.DataModel{
				{
					PluginInfo: model.PluginInfo{
						Events: []model.Events{
							{
								Id:             "1",
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
						Id: []string{"1"},
					},
				},
			},
			expected: RemoteRequestDTOItem{Events: []DTOEvents{
				{
					Id:             "1",
					CreatedAt:      "1",
					Type:           "1",
					Project:        "1",
					ProjectBaseDir: "some-base",
					Language:       "1",
					Target:         "1",
					Branch:         "master",
					Params:         nil,
				},
			}},
		},
		"some model": {
			models: []model.DataModel{
				{
					PluginInfo: model.PluginInfo{
						Events: []model.Events{
							{
								Id:             "1",
								CreatedAt:      "1",
								Type:           "1",
								Project:        "1",
								ProjectBaseDir: "some-base",
								Language:       "1",
								Target:         "1",
								Branch:         "",
								Params:         nil,
							},
							{
								Id:             "2",
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
					AggregatorInfo: model.AggregatorInfo{
						GitBranchesByProjectBaseDir: map[string]string{
							"some-base":  "master",
							"some-base2": "master2",
						},
						Id: []string{"1", "2"},
					},
				},
			},
			expected: RemoteRequestDTOItem{Events: []DTOEvents{
				{
					Id:             "1",
					CreatedAt:      "1",
					Type:           "1",
					Project:        "1",
					ProjectBaseDir: "some-base",
					Language:       "1",
					Target:         "1",
					Branch:         "master",
					Params:         nil,
				},
				{
					Id:             "2",
					CreatedAt:      "2",
					Type:           "2",
					Project:        "2",
					ProjectBaseDir: "some-base2",
					Language:       "2",
					Target:         "2",
					Branch:         "master2",
					Params:         nil,
				},
			}},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewRemoteRequestDTOFromDataModels(tc.models))
		})
	}
}
