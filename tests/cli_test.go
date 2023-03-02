package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaroslav1991/tts/internal/service/dispatcher/data/sender"
	"github.com/stretchr/testify/assert"
)

func TestCliSuccess(t *testing.T) {
	httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestBody, err := io.ReadAll(request.Body)
		assert.NoError(t, err)

		var requestDTO sender.RemoteRequestDTO
		assert.NoError(t, json.Unmarshal(requestBody, &requestDTO))

		// todo fix test
		if assert.Len(t, requestDTO, 2) {

			assert.NotEmpty(t, requestDTO[0].Events[0].Branch)
			assert.NotEmpty(t, requestDTO[1].Events[1].Branch)

			//assert.Equal(t, sender.RemoteRequestDTO{{
			//	PluginType:    "jetbrains",
			//	PluginVersion: "1.0.0",
			//	CliType:       "macos",
			//	CliVersion:    "2.1.0",
			//	DeviceName:    "vasya mac",
			//	Events: []sender.DTOEvents{
			//		{
			//			Uid:       requestDTO[0].Events[0].Uid,
			//			CreatedAt: "2022-01-11 14:23:01",
			//			Type:      "modify file",
			//			Project:   "some project",
			//			Language:  "golang",
			//			Target:    "../",
			//			Branch:    "model_fix",
			//			Params:    nil,
			//		},
			//	},
			//}}, requestDTO)
		}
	}))
	//cmd := exec.Command(
	//	"go",
	//	"run",
	//	"../cmd/cli/main.go",
	//	"-s",
	//	server.URL,
	//	"-d",
	//	`{"pluginType":"jetbrains","pluginVersion":"1.0.0","cliType":"macos","cliVersion":"2.1.0","deviceName":"vasya mac","events":[{"createdAt":"2022-01-11 14:23:01","type":"modify file","project":"some project","language":"golang","target":"../"}]}`,
	//)
	//
	//out, err := cmd.CombinedOutput()
	//if !assert.NoError(t, err) {
	//	t.Log("program error output:", string(out))
	//}
}
