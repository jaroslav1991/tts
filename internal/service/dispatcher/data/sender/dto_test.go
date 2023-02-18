package sender

import (
	"testing"
	"time"

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
						Program:     "program",
						Duration:    time.Second,
						PathProject: "some project",
					},
					AggregatorInfo: model.AggregatorInfo{
						CurrentGitBranch: "master",
					},
				},
			},
			expected: RemoteRequestDTO{
				{
					Program:          "program",
					Duration:         time.Second,
					PathProject:      "some project",
					CurrentGitBranch: "master",
				},
			},
		},
		"some model": {
			models: []model.DataModel{
				{
					PluginInfo: model.PluginInfo{
						Program:     "program",
						Duration:    time.Second,
						PathProject: "some project",
					},
					AggregatorInfo: model.AggregatorInfo{
						CurrentGitBranch: "master",
					},
				},
				{
					PluginInfo: model.PluginInfo{
						Program:     "program 2",
						Duration:    2 * time.Second,
						PathProject: "some project 2",
					},
					AggregatorInfo: model.AggregatorInfo{
						CurrentGitBranch: "some-branch",
					},
				},
			},
			expected: RemoteRequestDTO{
				{
					Program:          "program",
					Duration:         time.Second,
					PathProject:      "some project",
					CurrentGitBranch: "master",
				},
				{
					Program:          "program 2",
					Duration:         2 * time.Second,
					PathProject:      "some project 2",
					CurrentGitBranch: "some-branch",
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
