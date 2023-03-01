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
		DeviceName:    nil,
		Events: model.Events{
			Uid:       "1",
			CreatedAt: "1",
			Type:      "1",
			Project:   nil,
			Language:  nil,
			Target:    nil,
			Branch:    nil,
			Params:    nil,
		},
	}
	master := "master"

	aggregationInfo := model.AggregatorInfo{
		CurrentGitBranch: &master,
	}

	actualData, err := preparer.PrepareData(pluginInfo, aggregationInfo)
	assert.NoError(t, err)

	expected := `{"PluginInfo":{"PluginType":"1","PluginVersion":"1","CliType":"1","CliVersion":"1","DeviceName":null,"Events":{"Uid":"1","CreatedAt":"1","Type":"1","Project":null,"Language":null,"Target":null,"Branch":null,"Params":null}},"AggregatorInfo":{"CurrentGitBranch":"master"}}`
	assert.Equal(t, expected, string(actualData))
}
