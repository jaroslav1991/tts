package aggregator

import (
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentBranchAggregator_Aggregate_BranchNotFound(t *testing.T) {
	branch := CurrentBranchAggregator{}
	expectedBranch := CurrentBranchAggregator{}
	info := model.PluginInfo{PathProject: ""}
	target := model.AggregatorInfo{CurrentGitBranch: "undefined"}

	actualErr := branch.Aggregate(info, &target)
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedBranch.Aggregate(info, &model.AggregatorInfo{CurrentGitBranch: "undefined"}), actualErr)
}
