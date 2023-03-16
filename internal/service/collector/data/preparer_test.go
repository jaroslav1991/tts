package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestPreparer_PrepareData(t *testing.T) {
	preparer := Preparer{}

	pluginInfo := model.PluginInfo{
		Uid:           "qwerty123",
		PluginType:    "1",
		PluginVersion: "1",
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
				Params:         nil,
			},
		},
	}

	aggregationInfo := model.AggregatorInfo{
		GitBranchesByProjectBaseDir: map[string]string{
			"1": "some-branch",
		},
	}

	actualData, err := preparer.PrepareData(pluginInfo, aggregationInfo)
	assert.NoError(t, err)

	expected := `{"PluginInfo":{"Uid":"qwerty123","PluginType":"1","PluginVersion":"1","IdeType":"","IdeVersion":"","Events":[{"CreatedAt":"1","Type":"1","Project":"","ProjectBaseDir":"1","Language":"","Target":"","Branch":"","Params":null}]},"AggregatorInfo":{"GitBranchesByProjectBaseDir":{"1":"some-branch"}}}`
	assert.Equal(t, expected, string(actualData))
}
