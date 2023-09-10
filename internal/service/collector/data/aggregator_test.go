package data

import (
	"errors"
	"testing"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAggregator_Aggregate_NoMergeAggregators(t *testing.T) {
	aggregator := Aggregator{}

	info := model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "1.1.0",
		OSName:        "",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				Id:             "",
				CreatedAt:      "1",
				Type:           "",
				Project:        "",
				ProjectBaseDir: "",
				Language:       "",
				Target:         "",
				Branch:         "",
				Timezone:       "",
				Params:         nil,
			},
		},
	}

	actualAggregatedInfo := model.AggregatorInfo{}

	actualData, err := aggregator.Aggregate(info)
	assert.NoError(t, err)

	assert.Equal(t, actualAggregatedInfo, actualData)
}

type errorMergeAggregatorMock struct{}

func (e errorMergeAggregatorMock) Aggregate(info model.PluginInfo, target *model.AggregatorInfo) error {
	return errors.New("some error")
}

func TestAggregator_Aggregate_Negative(t *testing.T) {
	aggregatorMock := errorMergeAggregatorMock{}

	aggregator := Aggregator{
		Aggregators: []MergeAggregator{aggregatorMock},
	}

	_, err := aggregator.Aggregate(model.PluginInfo{})
	assert.Error(t, err)
}

type mergeAggregatorMock struct{}

func (m mergeAggregatorMock) Aggregate(info model.PluginInfo, target *model.AggregatorInfo) error {
	target.GitBranchesByProjectBaseDir = map[string]string{
		"some-base": "some-branch",
	}

	target.OSName = "windows"

	target.Id = "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84"
	return nil
}

func TestAggregator_Aggregate_Positive(t *testing.T) {
	aggregatorMock := mergeAggregatorMock{}

	aggregator := Aggregator{
		Aggregators: []MergeAggregator{aggregatorMock},
	}

	info := model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "windowsOS",
		CliVersion:    "1.1.0",
		OSName:        "",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				Id:             "",
				CreatedAt:      "1",
				Type:           "1",
				Project:        "",
				ProjectBaseDir: "some-base",
				Language:       "",
				Target:         "",
				Branch:         "",
				Timezone:       "",
				Params:         nil,
			},
		},
	}

	expected := model.AggregatorInfo{
		GitBranchesByProjectBaseDir: map[string]string{
			"some-base": "some-branch",
		},
		OSName: "windows",
		Id:     "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
	}

	actualRes, err := aggregator.Aggregate(info)
	assert.NoError(t, err)
	assert.Equal(t, expected, actualRes)
}
