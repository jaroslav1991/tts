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
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    "",
		Events: []model.Events{
			{
				Uid:       "some-uuid",
				CreatedAt: "1",
				Type:      "1",
				Project:   "",
				Language:  "",
				Target:    "",
				Branch:    "",
				Params:    nil,
			},
		},
	}

	aggregationInfo := model.AggregatorInfo{
		GitBranchesByEventUID: map[string]string{
			"some-uuid": "some-branch",
		},
	}

	actualData, err := preparer.PrepareData(pluginInfo, aggregationInfo)
	assert.NoError(t, err)

	expected := `{"PluginInfo":{"PluginType":"1","PluginVersion":"1","CliType":"1","CliVersion":"1","DeviceName":"","Events":[{"Uid":"some-uuid","CreatedAt":"1","Type":"1","Project":"","Language":"","Target":"","Branch":"","Params":null}]},"AggregatorInfo":{"GitBranchesByEventUID":{"some-uuid":"some-branch"}}}`
	assert.Equal(t, expected, string(actualData))
}
