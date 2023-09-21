package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestPreparer_PrepareData(t *testing.T) {
	preparer := Preparer{}

	pluginInfo := model.PluginInfo{
		Events: []model.Events{
			{
				Id:             "483a2111-575d-4a59-802c-d71246680bcb",
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
	}

	actualData, err := preparer.PrepareData(pluginInfo, aggregationInfo)
	assert.NoError(t, err)

	expected := `{"PluginInfo":{"Events":[{"Id":"483a2111-575d-4a59-802c-d71246680bcb","CreatedAt":"1","Type":"1","Project":"","ProjectBaseDir":"1","Language":"","Target":"","Branch":"","Timezone":"","Params":null}]},"AggregatorInfo":{"GitBranchesByProjectBaseDir":{"1":"some-branch"},"Id":null}}`
	assert.Equal(t, expected, string(actualData))
}
