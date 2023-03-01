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
		CliVersion:    "1",
		DeviceName:    "",
		Events: []model.Events{
			{
				Uid:       "qwerty",
				CreatedAt: "1",
				Type:      "1",
				Project:   "",
				Language:  "",
				Target:    "",
				Branch:    "",
				Params:    nil,
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
	target.GitBranchesByEventUID = map[string]string{
		"some-uid": "some-branch",
	}
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
		DeviceName:    "",
		Events: []model.Events{
			{
				Uid:       "some-uid",
				CreatedAt: "1",
				Type:      "1",
				Project:   "",
				Language:  "",
				Target:    "",
				Branch:    "",
				Params:    nil,
			},
		},
	}

	expected := model.AggregatorInfo{
		GitBranchesByEventUID: map[string]string{
			"some-uid": "some-branch",
		},
	}

	actualRes, err := aggregator.Aggregate(info)
	assert.NoError(t, err)
	assert.Equal(t, expected, actualRes)
}
