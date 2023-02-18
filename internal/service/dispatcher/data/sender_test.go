package data

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/model"
)

func TestSender_Send_Positive(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		assert.NoError(t, err)
		assert.Equal(t, `[{"Program":"test1","Duration":2}]`, string(body))
	}))

	actualData := []model.PluginInfo{{Program: "test1", Duration: 2}}

	sender := Sender{HttpAddr: server.URL}
	actualErr := sender.Send(actualData)
	assert.NoError(t, actualErr)
}
