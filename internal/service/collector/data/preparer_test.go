package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestPreparer_PrepareData(t *testing.T) {
	preparer := Preparer{}

	pluginInfo := model.PluginInfo{
		Program:  "test1",
		Duration: 0,
	}

	aggregationInfo := model.AggregatorInfo{
		CurrentGitBranch: "master",
	}

	actualData, err := preparer.PrepareData(pluginInfo, aggregationInfo)
	assert.NoError(t, err)

	expected := `{"PluginInfo":{"Program":"test1","Duration":0},"AggregatorInfo":{"CurrentGitBranch":"master"}}`
	assert.Equal(t, expected, string(actualData))
}
