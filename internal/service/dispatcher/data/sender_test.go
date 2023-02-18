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
		assert.Equal(t, `[{"PluginInfo":{"Program":"test","Duration":5,"PathProject":"testPath"},"AggregatorInfo":{"CurrentGitBranch":"testBranch"}}]`, string(body))
	}))

	actualData := []model.DataModel{{PluginInfo: model.PluginInfo{Program: "test", Duration: 5, PathProject: "testPath"}, AggregatorInfo: model.AggregatorInfo{CurrentGitBranch: "testBranch"}}}

	sender := Sender{HttpAddr: server.URL}
	actualErr := sender.Send(actualData)
	assert.NoError(t, actualErr)
}
