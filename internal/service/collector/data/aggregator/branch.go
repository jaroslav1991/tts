package aggregator

import (
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
)

type CurrentBranchAggregator struct {
	data.MergeAggregator
}

func (a *CurrentBranchAggregator) Aggregate(
	info model.PluginInfo,
	target *model.AggregatorInfo,
) error {
	// todo get branch for project from pluginInfo
	target.CurrentGitBranch = "master"
	return nil
}
