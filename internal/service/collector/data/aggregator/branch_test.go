package aggregator

import (
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentBranchAggregator_Aggregate_BranchNotFound(t *testing.T) {
	branch := CurrentBranchAggregator{}
	info := model.PluginInfo{}
	info.Events.Branch = nil
	target := model.AggregatorInfo{CurrentGitBranch: "master"}

	actualErr := branch.Aggregate(info, &target)
	assert.NoError(t, actualErr)
}

func TestCurrentBranchAggregator_Aggregate_BranchFound(t *testing.T) {
	info := model.PluginInfo{}
	master := "master"
	info.Events.Branch = &master
	target := model.AggregatorInfo{}

	branch := CurrentBranchAggregator{}
	actualErr := branch.Aggregate(info, &target)
	assert.NoError(t, actualErr)
}
