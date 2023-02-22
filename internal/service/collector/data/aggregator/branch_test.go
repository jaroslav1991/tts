package aggregator

import (
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentBranchAggregator_Aggregate_Positive(t *testing.T) {
	branch := CurrentBranchAggregator{}
	info := model.PluginInfo{PathProject: "C:\\Users\\jaros\\GolandProjects\\tts"}
	target := model.AggregatorInfo{CurrentGitBranch: "aggregator-01"}

	actualErr := branch.Aggregate(info, &target)
	assert.NoError(t, actualErr)
}

func TestCurrentBranchAggregator_Aggregate_BranchNotFound(t *testing.T) {
	branch := CurrentBranchAggregator{}
	info := model.PluginInfo{PathProject: ""}
	target := model.AggregatorInfo{CurrentGitBranch: "undefined"}

	actualErr := branch.Aggregate(info, &target)
	assert.NoError(t, actualErr)
}
