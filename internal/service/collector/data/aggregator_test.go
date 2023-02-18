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
		Program:     "test",
		Duration:    5,
		PathProject: "testPath",
	}
	target := model.AggregatorInfo{
		CurrentGitBranch: "",
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
	target.CurrentGitBranch = "testBranch"
	return nil
}

func TestAggregator_Aggregate_Positive(t *testing.T) {
	aggregatorMock := mergeAggregatorMock{}

	aggregator := Aggregator{
		Aggregators: []MergeAggregator{aggregatorMock},
	}

	info := model.PluginInfo{
		Program:     "test",
		Duration:    5,
		PathProject: "testPath",
	}

	expected := model.AggregatorInfo{
		CurrentGitBranch: "testBranch",
	}

	actualRes, err := aggregator.Aggregate(info)
	assert.NoError(t, err)
	assert.Equal(t, expected, actualRes)
}
