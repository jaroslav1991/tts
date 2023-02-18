package http

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestDataReader_ReadData_Positive(t *testing.T) {
	requestData := []byte(`{
		"program": "Some IDE",
		"durationMS": 15000
	}`)

	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestData))
	reader := DataReader{}

	actualData, actualErr := reader.ReadData(request)
	assert.NoError(t, actualErr)

	assert.Equal(t, model.PluginInfo{
		Program:  "Some IDE",
		Duration: 15 * time.Second,
	}, actualData)
}

func TestDataReader_ReadData_Negative_UnmarshalFailed(t *testing.T) {
	requestData := []byte(`123`)

	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestData))
	reader := DataReader{}

	_, actualErr := reader.ReadData(request)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, ErrUnmarshalRequestData)
}

type errReader struct {
}

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("some read error")
}

func TestDataReader_ReadData_Negative_ReadDataError(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/", &errReader{})
	reader := DataReader{}

	_, actualErr := reader.ReadData(request)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, ErrReadBodyFailed)
}

func TestDataReader_ReadData_Negative_InvalidMethod(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	reader := DataReader{}

	_, actualErr := reader.ReadData(request)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, ErrInvalidRequestMethod)
}

func TestDataReader_ReadData_Negative_InvalidRequestType(t *testing.T) {
	reader := DataReader{}

	_, actualErr := reader.ReadData("banana")
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, ErrInvalidRequestType)
}
