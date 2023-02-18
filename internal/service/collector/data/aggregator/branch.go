package aggregator

import (
	"fmt"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
	"os"
	"strings"
)

type CurrentBranchAggregator struct {
	data.MergeAggregator
}

func (a *CurrentBranchAggregator) Aggregate(
	info model.PluginInfo,
	target *model.AggregatorInfo,
) error {
	path := info.PathProject

	currentBranch, err := os.ReadFile(path + string(os.PathSeparator) + ".git" + string(os.PathSeparator) + "HEAD")
	if err != nil {
		return fmt.Errorf("reading .git\\HEAD failed: %w", err)
	}
	target.CurrentGitBranch = strings.TrimSpace(strings.ReplaceAll(string(currentBranch), "ref: refs/heads/", ""))

	return nil
}
