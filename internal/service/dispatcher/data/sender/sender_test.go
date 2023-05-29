package sender

import (
	"bytes"
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSender_Send_Positive(t *testing.T) {
	reqData := `[{"uid":"qwerty123","pluginType":"1","pluginVersion":"1","ideType":"1","ideVersion":"1","events":[{"createdAt":"1","type":"1","project":"1","projectBaseDir":"some-base","language":"1","target":"1","branch":"some-branch"}]}]`

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		assert.NoError(t, err)
		assert.Equal(t, reqData, string(body))
	}))

	req := httptest.NewRequest("POST", "/events", bytes.NewBuffer([]byte(reqData)))
	req.Header.Set("Authorization", "token")

	actualData := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				Uid:           "qwerty123",
				PluginType:    "1",
				PluginVersion: "1",
				IdeType:       "1",
				IdeVersion:    "1",
				Events: []model.Events{
					{
						CreatedAt:      "1",
						Type:           "1",
						Project:        "1",
						ProjectBaseDir: "some-base",
						Language:       "1",
						Target:         "1",
						Branch:         "some-branch",
						Params:         nil,
					},
				},
			},
			AggregatorInfo: model.AggregatorInfo{
				GitBranchesByProjectBaseDir: map[string]string{
					"some-base": "some-branch",
				},
			},
		},
	}

	sender := Sender{HttpAddr: server.URL, AuthKey: "token"}
	actualErr := sender.Send(actualData)
	assert.NoError(t, actualErr)
	assert.Equal(t, "token", req.Header.Get("Authorization"))
}
