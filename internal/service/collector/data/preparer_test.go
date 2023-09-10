package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestPreparer_PrepareData(t *testing.T) {
	preparer := Preparer{}

	pluginInfo := model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "windowsOS",
		CliVersion:    "1.1.0",
		IdeType:       "",
		IdeVersion:    "",
		Events: []model.Events{
			{
				CreatedAt:      "1",
				Type:           "1",
				Project:        "",
				ProjectBaseDir: "1",
				Language:       "",
				Target:         "",
				Branch:         "",
				Timezone:       "",
				Params:         nil,
			},
		},
	}

	aggregationInfo := model.AggregatorInfo{
		GitBranchesByProjectBaseDir: map[string]string{
			"1": "some-branch",
		},
		OSName: "windows",
		Id:     "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
	}

	actualData, err := preparer.PrepareData(pluginInfo, aggregationInfo)
	assert.NoError(t, err)

	expected := `{"PluginInfo":{"PluginType":"1","PluginVersion":"1","CliType":"windowsOS","CliVersion":"1.1.0","OSName":"","IdeType":"","IdeVersion":"","Events":[{"Id":"","CreatedAt":"1","Type":"1","Project":"","ProjectBaseDir":"1","Language":"","Target":"","Branch":"","Timezone":"","Params":null}]},"AggregatorInfo":{"GitBranchesByProjectBaseDir":{"1":"some-branch"},"OSName":"windows","Id":"a6ac8ef0-28e2-4b6e-8568-aa8934f53c84"}}`
	assert.Equal(t, expected, string(actualData))
}
