package aggregator

import (
	"github.com/google/uuid"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
	"log"
	"os"
	"strings"
)

var (
	getBranchFn = GetBranchByProjectBaseDir
	getIDFn     = GetUUID
)

type CommonAggregator struct {
	data.MergeAggregator
}

func (a *CommonAggregator) Aggregate(info model.PluginInfo, target *model.AggregatorInfo) error {
	target.GitBranchesByProjectBaseDir = map[string]string{}
	target.Id = make([]string, 0, len(info.Events))

	for i := range info.Events {

		if info.Events[i].Branch != "" {
			continue
		}

		if eventBranch := getBranchFn(info.Events[i].ProjectBaseDir); eventBranch != "" {
			target.GitBranchesByProjectBaseDir[info.Events[i].ProjectBaseDir] = eventBranch
		}

		if info.Events[i].Id != "" {
			target.Id = append(target.Id, info.Events[i].Id)
		} else if info.Events[i].Id == "" {
			target.Id = append(target.Id, getIDFn())
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

func GetUUID() string {
	id := uuid.New()

	return id.String()
}
