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
	reqData := `{"events":[{"id":"qwerty","createdAt":"2022-01-1114:23:01","type":"modifyfile","project":"someproject","projectBaseDir":"/mnt/c/Users/jaros/GolandProjects/tts","language":"golang","target":"C/Projects/Golang/cli-tts","branch":"new_contract_v1"}]}`

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		assert.NoError(t, err)
		assert.Equal(t, reqData, string(body))
		assert.Equal(t, "token", request.Header.Get("Authorization"))
	}))

	actualData := []model.DataModel{
		{
			PluginInfo: model.PluginInfo{
				Events: []model.Events{
					{
						Id:             "qwerty",
						CreatedAt:      "2022-01-1114:23:01",
						Type:           "modifyfile",
						Project:        "someproject",
						ProjectBaseDir: "/mnt/c/Users/jaros/GolandProjects/tts",
						Language:       "golang",
						Target:         "C/Projects/Golang/cli-tts",
						Branch:         "",
						Timezone:       "",
						Params:         nil,
					},
				},
			},
			AggregatorInfo: model.AggregatorInfo{
				GitBranchesByProjectBaseDir: map[string]string{
					"/mnt/c/Users/jaros/GolandProjects/tts": "new_contract_v1",
				},
			},
		},
	}

	sender := Sender{HttpAddr: server.URL, AuthKey: "token"}
	actualErr := sender.Send(actualData)
	assert.NoError(t, actualErr)
}
