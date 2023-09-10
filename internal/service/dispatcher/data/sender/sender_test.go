package sender

import (
	"github.com/jaroslav1991/tts/internal/model"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSender_Send_Positive(t *testing.T) {
	reqData := `[{"pluginType":"1","pluginVersion":"1","cliType":"windowsOS","cliVersion":"1.1.0","osName":"windows","ideType":"1","ideVersion":"1","events":[{"id":"a6ac8ef0-28e2-4b6e-8568-aa8934f53c84","createdAt":"1","type":"1","project":"1","projectBaseDir":"some-base","language":"1","target":"1","branch":"some-branch","timezone":"1"}]}]`

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		assert.NoError(t, err)
		assert.Equal(t, reqData, string(body))
		assert.Equal(t, "token", request.Header.Get("Authorization"))
	}))

	actualData := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				PluginType:    "1",
				PluginVersion: "1",
				CliType:       "windowsOS",
				CliVersion:    "1.1.0",
				OSName:        "windows",
				IdeType:       "1",
				IdeVersion:    "1",
				Events: []model.Events{
					{
						Id:             "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
						CreatedAt:      "1",
						Type:           "1",
						Project:        "1",
						ProjectBaseDir: "some-base",
						Language:       "1",
						Target:         "1",
						Branch:         "some-branch",
						Timezone:       "1",
						Params:         nil,
					},
				},
			},
			AggregatorInfo: model.AggregatorInfo{
				GitBranchesByProjectBaseDir: map[string]string{
					"some-base": "some-branch",
				},
				OSName: "windows",
				Id:     "a6ac8ef0-28e2-4b6e-8568-aa8934f53c84",
			},
		},
	}

	sender := Sender{HttpAddr: server.URL, AuthKey: "token"}
	actualErr := sender.Send(actualData)
	assert.NoError(t, actualErr)
}
