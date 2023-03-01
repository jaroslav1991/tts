package data

import (
	"errors"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregator_Aggregate_Empty(t *testing.T) {
	aggregator := Aggregator{}

	info := model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    "",
		Events: model.Events{
			Uid:       "qwerty",
			CreatedAt: "1",
			Type:      "1",
			Project:   "",
			Language:  "",
			Target:    "",
			Branch:    "",
			Params:    "",
		},
	}
	target := model.AggregatorInfo{
		CurrentGitBranch: nil,
	}

	actualData, err := aggregator.Aggregate(info)
	assert.NoError(t, err)

	assert.Equal(t, target, actualData)
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
	target.CurrentGitBranch = nil
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
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "qwerty",
			CreatedAt: "1",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	}

	expected := model.AggregatorInfo{
		CurrentGitBranch: nil,
	}

	actualRes, err := aggregator.Aggregate(info)
	assert.NoError(t, err)
	assert.Equal(t, expected, actualRes)
}
