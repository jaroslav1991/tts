package aggregator

import (
	"testing"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCurrentBranchAggregator_Aggregate_BranchNotFoundInEvent(t *testing.T) {
	aggregator := CurrentBranchAggregator{}
	pluginInfo := model.PluginInfo{
		Events: []model.Events{
			{
				Uid:    "some-uid",
				Branch: "",
			},
		},
	}

	assert.NoError(t, aggregator.Aggregate(pluginInfo, &model.AggregatorInfo{}))
}

func TestCurrentBranchAggregator_Aggregate_BranchNotFoundInEventAndFoundInGit(t *testing.T) {
	getBranchFn = func(target string) string {
		if target == "target-1" {
			return "some-branch-1"
		}

		if target == "target-2" {
			return "some-branch-2"
		}

		t.Errorf("unexpected target: %s", target)
		return ""
	}

	defer func() {
		getBranchFn = GetBranchByTarget
	}()

	aggregator := CurrentBranchAggregator{}
	pluginInfo := model.PluginInfo{
		Events: []model.Events{
			{
				Uid:    "some-uid-1",
				Branch: "",
				Target: "target-1",
			},
			{
				Uid:    "some-uid-2",
				Branch: "",
				Target: "target-2",
			},
		},
	}

	target := model.AggregatorInfo{}

	assert.NoError(t, aggregator.Aggregate(pluginInfo, &target))

	assert.Equal(t, model.AggregatorInfo{
		GitBranchesByEventUID: map[string]string{
			"some-uid-1": "some-branch-1",
			"some-uid-2": "some-branch-2",
		},
	}, target)

}

func TestCurrentBranchAggregator_Aggregate_BranchFoundInEvent(t *testing.T) {
	aggregator := CurrentBranchAggregator{}
	pluginInfo := model.PluginInfo{
		Events: []model.Events{
			{
				Uid:    "some-uid",
				Branch: "master",
			},
		},
	}

	target := model.AggregatorInfo{}

	actualErr := aggregator.Aggregate(pluginInfo, &target)
	assert.NoError(t, actualErr)
	assert.Empty(t, target.GitBranchesByEventUID)
}
