package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaroslav1991/tts/internal/service/dispatcher/data/sender"
)

func TestCliSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestBody, err := io.ReadAll(request.Body)
		assert.NoError(t, err)

		var requestDTO sender.RemoteRequestDTO
		assert.NoError(t, json.Unmarshal(requestBody, &requestDTO))

		if assert.Len(t, requestDTO, 1) {
			assert.NotEmpty(t, requestDTO[0].Events[0].Uid)
			assert.NotEmpty(t, requestDTO[0].Events[0].Branch)

			requestDTO[0].Events[0].Uid = "3607bbe0-2c9a-4c51-b636-5e6a7db8b574"
			requestDTO[0].Events[0].Branch = "some-branch"

			assert.Equal(t, sender.RemoteRequestDTO{{
				PluginType:    "jetbrains",
				PluginVersion: "1.0.0",
				CliType:       "macos",
				CliVersion:    "2.1.0",
				DeviceName:    "vasya mac",
				Events: []sender.DTOEvents{
					{
						Uid:       "3607bbe0-2c9a-4c51-b636-5e6a7db8b574",
						CreatedAt: "2022-01-11 14:23:01",
						Type:      "modify file",
						Project:   "some project",
						Language:  "golang",
						Target:    "../",
						Branch:    "some-branch",
					},
				},
			}}, requestDTO)
		}
	}))
	cmd := exec.Command(
		"go",
		"run",
		"../cmd/cli/main.go",
		"-s",
		server.URL,
		"-d",
		`{"pluginType":"jetbrains","pluginVersion":"1.0.0","cliType":"macos","cliVersion":"2.1.0","deviceName":"vasya mac","events":[{"createdAt":"2022-01-11 14:23:01","type":"modify file","project":"some project","language":"golang","target":"../"}]}`,
	)

	out, err := cmd.CombinedOutput()
	if !assert.NoError(t, err) {
		t.Log("program error output:", string(out))
	}
}
