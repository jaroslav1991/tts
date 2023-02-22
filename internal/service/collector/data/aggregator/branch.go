package aggregator

import (
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
	"log"
	"os"
	"strings"
)

type CurrentBranchAggregator struct {
	data.MergeAggregator
}

// todo implement if user not using "git" and write test

func (a *CurrentBranchAggregator) Aggregate(
	info model.PluginInfo,
	target *model.AggregatorInfo,
) error {
	path := info.PathProject

	currentBranch, err := os.ReadFile(path + string(os.PathSeparator) + ".git" + string(os.PathSeparator) + "HEAD")
	if err != nil {
		log.Printf("current branch path not found: %v", err)
		target.CurrentGitBranch = "undefined"
		return nil
	}
	target.CurrentGitBranch = strings.TrimSpace(strings.ReplaceAll(string(currentBranch), "ref: refs/heads/", ""))

	return nil
}
