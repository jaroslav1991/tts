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
		Uid:           "qwerty123",
		PluginType:    "1",
		PluginVersion: "1",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "1",
				Type:           "",
				Project:        "",
				ProjectBaseDir: "",
				Language:       "",
				Target:         "",
				Branch:         "",
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
	return nil
}

func TestAggregator_Aggregate_Positive(t *testing.T) {
	aggregatorMock := mergeAggregatorMock{}

	aggregator := Aggregator{
		Aggregators: []MergeAggregator{aggregatorMock},
	}

	info := model.PluginInfo{
		Uid:           "qwerty123",
		PluginType:    "1",
		PluginVersion: "1",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "1",
				Type:           "1",
				Project:        "",
				ProjectBaseDir: "some-base",
				Language:       "",
				Target:         "",
				Branch:         "",
				Params:         nil,
			},
		},
	}

	expected := model.AggregatorInfo{
		GitBranchesByProjectBaseDir: map[string]string{
			"some-base": "some-branch",
		},
	}

	actualRes, err := aggregator.Aggregate(info)
	assert.NoError(t, err)
	assert.Equal(t, expected, actualRes)
}
