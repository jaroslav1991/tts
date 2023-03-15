package aggregator

import (
	"log"
	"os"
	"strings"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
)

var getBranchFn = GetBranchByProjectBaseDir

type CurrentBranchAggregator struct {
	data.MergeAggregator
}

func (a *CurrentBranchAggregator) Aggregate(
	info model.PluginInfo,
	target *model.AggregatorInfo,
) error {
	target.GitBranchesByProjectBaseDir = map[string]string{}

	for i := range info.Events {
		if info.Events[i].Branch != "" {
			continue
		}

		if eventBranch := getBranchFn(info.Events[i].ProjectBaseDir); eventBranch != "" {
			//target.GitBranchesByEventUID[info.Events[i].Uid] = eventBranch
			target.GitBranchesByProjectBaseDir[info.Events[i].ProjectBaseDir] = eventBranch
		}
	}

	return nil
}

func GetBranchByProjectBaseDir(projectBaseDir string) string {
	filename := projectBaseDir + string(os.PathSeparator) + ".git" + string(os.PathSeparator) + "HEAD"

	currentBranch, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("current branch path not found: %v", err)
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(string(currentBranch), "ref: refs/heads/", ""))
}
