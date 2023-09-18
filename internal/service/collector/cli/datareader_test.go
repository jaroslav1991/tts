package cli

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func Test(t *testing.T) {
	b, _ := json.Marshal(DTO{
		Events: []DTOEvents{
			{
				Id:             "1a03028f-6f1c-43f9-a08d-c13c28fa97cb",
				CreatedAt:      "2022-01-11 14:23:01",
				Type:           "modify file",
				Project:        "some project",
				ProjectBaseDir: "some dir",
				Language:       "golang",
				Target:         "../",
				Timezone:       "123456789",
			},
		},
	})

	fmt.Println(string(b))
}

func TestDataReader_ReadData_Positive(t *testing.T) {
	requestData := `{
			"events":
				[
					{
						"id":"1a03028f-6f1c-43f9-a08d-c13c28fa97cb",
						"createdAt":"1",
						"type":"1",
						"project":"1",
						"language":"1",
						"target":"1",
						"branch":"1",
						"timezone":"123456789"
					}
				]
			}`

	reader := DataReader{}

	actualData, actualErr := reader.ReadData(requestData)
	assert.NoError(t, actualErr)

	assert.Equal(t, model.PluginInfo{
		Events: []model.Events{
			{
				Id:        "1a03028f-6f1c-43f9-a08d-c13c28fa97cb",
				CreatedAt: "1",
				Type:      "1",
				Project:   "1",
				Language:  "1",
				Target:    "1",
				Branch:    "1",
				Timezone:  "123456789",
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
