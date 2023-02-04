package cli

import (
	"github.com/jaroslav1991/tts/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDataReader_ReadData_Positive(t *testing.T) {
	requestData := `{
		"program": "Some IDE",
		"durationMS": 15000
	}`

	reader := DataReader{}

	actualData, actualErr := reader.ReadData(requestData)
	assert.NoError(t, actualErr)

	assert.Equal(t, service.DataModel{
		Program:  "Some IDE",
		Duration: 15 * time.Second,
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
