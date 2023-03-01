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
	for i := range info.Events {
		if info.Events[i].Branch == "" {
			info.Events[i].Branch = GetBranchByTarget(info.Events[i].Target)
			info.Events[i].Branch = target.CurrentGitBranch
		}
	}

	return nil
}

func GetBranchByTarget(target string) string {
	filename := target + string(os.PathSeparator) + ".git" + string(os.PathSeparator) + "HEAD"

	currentBranch, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("current branch path not found: %v", err)
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(string(currentBranch), "ref: refs/heads/", ""))
}
