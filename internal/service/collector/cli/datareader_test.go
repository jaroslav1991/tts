package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestDataReader_ReadData_Positive(t *testing.T) {
	requestData := `{
			"pluginType":"1",
			"pluginVersion":"1",
			"cliType":"1",
			"cliVersion":"1",
			"deviceName":"1",
			"events":
				[
					{
						"uid":"1",
						"createdAt":"1",
						"type":"1",
						"project":"1",
						"language":"1",
						"target":"1",
						"branch":"1"
					}
				]
			}`

	reader := DataReader{}

	actualData, actualErr := reader.ReadData(requestData)
	assert.NoError(t, actualErr)

	assert.Equal(t, model.PluginInfo{
		PluginType:    "1",
		PluginVersion: "1",
		CliType:       "1",
		CliVersion:    "1",
		DeviceName:    "1",
		Events: []model.Events{
			{
				Uid:       "1",
				CreatedAt: "1",
				Type:      "1",
				Project:   "1",
				Language:  "1",
				Target:    "1",
				Branch:    "1",
				Params:    nil,
			},
		},
	}, actualData)
}

func TestDataReader_ReadData_Negative_UnmarshalFailed(t *testing.T) {
	requestData := `{qwertyui}`

	reader := DataReader{}

	_, actualErr := reader.ReadData(requestData)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, ErrUnmarshalRequestData)
}

func TestDataReader_ReadData_Negative_InvalidRequestType(t *testing.T) {
	reader := DataReader{}

	_, actualErr := reader.ReadData(123)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, ErrInvalidRequestType)
}
