package aggregator

import (
	"testing"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCommonAggregator_Aggregate_BranchNotFoundInEvent(t *testing.T) {
	aggregator := CommonAggregator{}
	pluginInfo := model.PluginInfo{
		Events: []model.Events{
			{
				Branch: "",
			},
		},
	}

	assert.NoError(t, aggregator.Aggregate(pluginInfo, &model.AggregatorInfo{}))
}

func TestCommonAggregator_Aggregate_BranchNotFoundInEventAndFoundInGit(t *testing.T) {
	getBranchFn = func(projectBaseDir string) string {
		if projectBaseDir == "some-base-1" {
			return "some-branch-1"
		}

		if projectBaseDir == "some-base-2" {
			return "some-branch-2"
		}

		t.Errorf("unexpected projectBaseDir: %s", projectBaseDir)
		return ""
	}

	defer func() {
		getBranchFn = GetBranchByProjectBaseDir
	}()

	aggregator := CommonAggregator{}
	pluginInfo := model.PluginInfo{
		Events: []model.Events{
			{
				Branch:         "",
				ProjectBaseDir: "some-base-1",
			},
			{
				Branch:         "",
				ProjectBaseDir: "some-base-2",
			},
		},
	}

	target := model.AggregatorInfo{}

	assert.NoError(t, aggregator.Aggregate(pluginInfo, &target))

	assert.Equal(t, model.AggregatorInfo{
		GitBranchesByProjectBaseDir: map[string]string{
			"some-base-1": "some-branch-1",
			"some-base-2": "some-branch-2",
		},
		Id: target.Id,
	}, target)

}

func TestCommonAggregator_Aggregate_BranchFoundInEvent(t *testing.T) {
	aggregator := CommonAggregator{}
	pluginInfo := model.PluginInfo{
		Events: []model.Events{
			{
				Branch: "master",
			},
		},
	}

	target := model.AggregatorInfo{}

	actualErr := aggregator.Aggregate(pluginInfo, &target)
	assert.NoError(t, actualErr)
	assert.Empty(t, target.GitBranchesByProjectBaseDir)
}

func TestCommonAggregate_IDNotEmpty(t *testing.T) {
	aggregator := CommonAggregator{}
	pluginInfo := model.PluginInfo{
		Events: []model.Events{
			{
				Id: "qwerty",
			},
		},
	}

	target := model.AggregatorInfo{
		GitBranchesByProjectBaseDir: nil,
	}

	assert.NoError(t, aggregator.Aggregate(pluginInfo, &target))
	assert.Equal(t, model.AggregatorInfo{GitBranchesByProjectBaseDir: map[string]string{}}, target)
}
