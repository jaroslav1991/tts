package aggregator

import (
	"github.com/google/uuid"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/jaroslav1991/tts/internal/model"
	"github.com/jaroslav1991/tts/internal/service/collector/data"
)

var (
	getBranchFn = GetBranchByProjectBaseDir
	currentOS   = GetOSName
	uid         = GetUUID
)

type CommonAggregator struct {
	data.MergeAggregator
}

func (a *CommonAggregator) Aggregate(info model.PluginInfo, target *model.AggregatorInfo) error {
	target.GitBranchesByProjectBaseDir = map[string]string{}

	for i := range info.Events {
		if info.Events[i].Branch != "" {
			continue
		}

		if eventBranch := getBranchFn(info.Events[i].ProjectBaseDir); eventBranch != "" {
			target.GitBranchesByProjectBaseDir[info.Events[i].ProjectBaseDir] = eventBranch
		}

		if info.Events[i].Uid == "" {
			target.Uid = uid()
		}
	}

	if info.OSName == "" {
		if osName := currentOS(); osName != "" {
			target.OSName = osName
		}
	}

	return nil
}

func GetOSName() string {
	return runtime.GOOS
}

func GetUUID() string {
	uid := uuid.New()

	return uid.String()
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
